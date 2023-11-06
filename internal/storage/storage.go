package storage

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/kudagonbe/s3-static-api/internal/config"
)

//go:embed static/*
var files embed.FS

func GetObject(key string) ([]byte, error) {
	cfg := config.Get()
	params := &s3.GetObjectInput{
		Bucket: aws.String(cfg.Bucket),
		Key:    aws.String(key),
	}
	resp, err := s3.NewFromConfig(cfg.AwsConfig, func(o *s3.Options) {
		o.UsePathStyle = true
	}).GetObject(context.Background(), params, s3.WithAPIOptions(
		v4.SwapComputePayloadSHA256ForUnsignedPayloadMiddleware,
	))
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	defer resp.Body.Close()
	return buf.Bytes(), nil
}

func PutObject(key string, addTimeStamp bool) error {
	b, e := files.ReadFile(fmt.Sprintf("static/%s", key))
	if e != nil {
		return e
	}
	cfg := config.Get()
	params := &s3.PutObjectInput{
		Bucket:        aws.String(cfg.Bucket),
		Key:           aws.String(generateObjectID(key, addTimeStamp)),
		Body:          bytes.NewBuffer(b),
		ContentLength: int64(len(b)),
	}
	_, err := s3.NewFromConfig(cfg.AwsConfig, func(o *s3.Options) {
		o.UsePathStyle = true
	}).PutObject(context.Background(), params, s3.WithAPIOptions(
		v4.SwapComputePayloadSHA256ForUnsignedPayloadMiddleware,
	))
	if err != nil {
		return err
	}
	return nil
}

func generateObjectID(key string, addTimeStamp bool) string {
	if addTimeStamp {
		t := time.Now()
		return fmt.Sprintf("%s%09d-%s", t.Format("20060102150405"), t.Nanosecond(), key)
	}
	return key
}
