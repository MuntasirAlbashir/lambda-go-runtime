package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io/ioutil"
	"log"
)

var (
	outBucketName = "golang-output-bucket"
)

type Response struct {
	Message string `json:"message:"`
}

var cfg aws.Config
var s3Svc *s3.Client

func init() {
	cfg, _ = config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	s3Svc = s3.NewFromConfig(cfg)
}

type S3GetObjectApi interface {
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(options *s3.Options)) (*s3.GetObjectOutput, error)
}

type S3PutObjectApi interface {
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(options *s3.Options)) (*s3.PutObjectOutput, error)
}

func UploadObject(objectName *string, object *[]byte, s3Api S3PutObjectApi) (err error) {
	log.Printf("uploading object %v\n with data %v\n", objectName, bytes.NewReader(*object))
	if object == nil || objectName == nil {
		log.Panicf("Ran into an error retrieving object %v", err)
	}
	res, err := s3Api.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(outBucketName),
		Key:    objectName,
		Body:   bytes.NewReader(*object),
	})

	if err != nil {
		log.Panicf("ran into the error %v", err)
		return err
	}
	log.Printf("object successfully uploaded %+v", res)
	return nil
}

func GetObject(bucketName, objectName *string, s3Api S3GetObjectApi) (object *s3.GetObjectOutput, err error) {
	object, err = s3Api.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: bucketName,
		Key:    objectName,
	})
	if err != nil {
		log.Printf("Ran into an error processing the object %v", err)
		return nil, err
	}
	log.Printf("successfully recieved object %v", object)

	return object, nil
}

func handler(event events.S3Event) (Response, error) {
	log.Printf("recieved event %+v", event)
	var bucketName string
	var objectName string

	for _, attr := range event.Records {
		bucketName = attr.S3.Bucket.Name
		objectName = attr.S3.Object.Key
	}

	obj, err := GetObject(&bucketName, &objectName, s3Svc)
	if err != nil {
		return Response{
			Message: err.Error(),
		}, err
	}

	data, err := ioutil.ReadAll(obj.Body)
	log.Printf("uploading object %v\n", string(data))
	if err != nil {
		log.Fatal(err)
	}
	err = UploadObject(&objectName, &data, s3Svc)

	if err != nil {
		return Response{
			Message: fmt.Sprintf("%v error recieved", err),
		}, err
	}
	return Response{
		Message: fmt.Sprintf("%+v successfully uploaded \n", objectName),
	}, nil
}

func main() {
	lambda.Start(handler)
}
