package main_test

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"testing"
)

type mockGetObject func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(options *s3.Options)) (*s3.GetObjectOutput, error)
type mockPutObject func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(options *s3.Options)) (*s3.GetObjectOutput, error)

func TestHandler(t *testing.T) {

}
