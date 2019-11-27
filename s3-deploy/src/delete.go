package main

import (
	"fmt"
	"math"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/maurerlabs/actions/toolkit"
)

func batchDelete(sess *session.Session, paths []*string, bucket, prefix string) error {
	total := len(paths)
	if total < 1 {
		toolkit.Info("Nothing to delete")
		return nil
	}

	word := "files"
	if total == 1 {
		word = "file"
	}

	toolkit.StartGroup(fmt.Sprintf("%v %s staged for deletion", total, word))
	defer toolkit.EndGroup()

	var recurse func(p []*string) error
	recurse = func(p []*string) error {
		max := 1000
		length := len(p)

		keys := math.Min(float64(length), float64(max))
		if keys < 1 {
			return nil
		}

		var objects []*s3.ObjectIdentifier
		for _, v := range p[:int(keys)] {
			key := prefix + *v
			objects = append(objects, &s3.ObjectIdentifier{
				Key: &key,
			})
		}

		res, err := s3.New(sess).DeleteObjects(&s3.DeleteObjectsInput{
			Bucket: aws.String(bucket),
			Delete: &s3.Delete{
				Quiet:   aws.Bool(false),
				Objects: objects,
			},
		})
		if err != nil {
			return err
		}

		for _, v := range res.Deleted {
			toolkit.Info(fmt.Sprintf("Deleted %s", *v.Key))
		}

		if length > max {
			return recurse(p[max:])
		}

		return nil
	}

	return recurse(paths)
}
