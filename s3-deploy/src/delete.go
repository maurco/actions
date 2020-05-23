package main

import (
	"fmt"
	"math"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/maurco/actions/toolkit"
)

func batchDelete(sess *session.Session, paths []*string, bucket, prefix string) error {
	total := len(paths)
	title := "Nothing to delete"

	if total > 0 {
		word := "files"
		if total == 1 {
			word = "file"
		}
		title = fmt.Sprintf("Deleting %v %s", total, word)
	}

	toolkit.StartGroup(title)
	defer toolkit.EndGroup()

	if total < 1 {
		return nil
	}

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
			toolkit.Info("Deleted %s", *v.Key)
		}

		if length > max {
			return recurse(p[max:])
		}

		return nil
	}

	return recurse(paths)
}
