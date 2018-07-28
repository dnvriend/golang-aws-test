package aws

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

func ListBuckets(svc *s3.S3) *s3.ListBucketsOutput {
	result, err := svc.ListBuckets(&s3.ListBucketsInput{})
	CheckErr(err)
	return result
}
