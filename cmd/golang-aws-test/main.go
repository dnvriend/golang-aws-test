package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dnvriend/golang-aws-test/internal/aws"
)

func main() {
	sess := session.Must(session.NewSession())
	svc := s3.New(sess)
	result := aws.ListBuckets(svc)
	fmt.Println(result)
}
