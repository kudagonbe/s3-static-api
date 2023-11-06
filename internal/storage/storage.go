package storage

import (
	"bytes"
	"context"
	"embed"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/kudagonbe/s3-static-api/internal/config"
)

//go:embed static/*
var files embed.FS

func PutObject(key string) error {
	b, e := files.ReadFile(fmt.Sprintf("static/%s", key))
	if e != nil {
		return e
	}
	cfg := config.Get()
	params := &s3.PutObjectInput{
		Bucket:        aws.String(cfg.Bucket),
		Key:           aws.String(key),
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
