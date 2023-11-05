package storage

import (
	"context"
	"net/url"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/kudagonbe/s3-static-api/internal/config"
)

func PutObject(key string) error {
	cfg := config.Get()
	params := &s3.PutObjectInput{
		Bucket: aws.String(cfg.Bucket),
		Key:    aws.String(key),
		Body:   strings.NewReader("Hello World!!"),
	}
	_, err := s3.NewFromConfig(cfg.AwsConfig, func(o *s3.Options) {
		o.EndpointResolverV2 = &resolverV2{}
	}).PutObject(context.Background(), params)
	if err != nil {
		return err
	}
	return nil
}

type resolverV2 struct{}

func (*resolverV2) ResolveEndpoint(ctx context.Context, params s3.EndpointParameters) (smithyendpoints.Endpoint, error) {
	cfg := config.Get()
	u, err := url.Parse(cfg.Endpoint)
	if err != nil {
		return smithyendpoints.Endpoint{}, err
	}
	return smithyendpoints.Endpoint{URI: *u}, nil
}
