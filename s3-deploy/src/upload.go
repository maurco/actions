package main

import (
	"fmt"
	"mime"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/dustin/go-humanize"
)

const UPLOAD_PART_SIZE = 1024 * 1024 * 10 // 10mb

type UploadIterator struct {
	bucket string
	prefix string
	acl    string
	paths  []*string
	err    error
	next   struct {
		path string
		f    *os.File
	}
}

func (iter *UploadIterator) Next() bool {
	if len(iter.paths) == 0 {
		iter.next.f = nil
		return false
	}

	f, err := os.Open(*iter.paths[0])
	iter.err = err

	iter.next.f = f
	iter.next.path = *iter.paths[0]
	iter.paths = iter.paths[1:]

	return iter.Err() == nil
}

func (iter *UploadIterator) Err() error {
	return iter.err
}

func (iter *UploadIterator) UploadObject() s3manager.BatchUploadObject {
	mimeType := mime.TypeByExtension(iter.next.path)
	key := iter.prefix + iter.next.path

	return s3manager.BatchUploadObject{
		Object: &s3manager.UploadInput{
			ACL:         &iter.acl,
			Bucket:      &iter.bucket,
			Key:         &key,
			Body:        iter.next.f,
			ContentType: &mimeType,
		},
		After: func() error {
			fi, err := iter.next.f.Stat()
			if err != nil {
				return err
			}

			size := humanize.Bytes(uint64(fi.Size()))
			fmt.Printf("-- Uploaded %s [%v]\n", key, size)
			return iter.next.f.Close()
		},
	}
}

func batchUpload(sess *session.Session, paths []*string, bucket, prefix, acl string) error {
	iter := &UploadIterator{
		bucket: bucket,
		prefix: prefix,
		acl:    acl,
		paths:  paths,
	}

	total := len(paths)

	if total < 1 {
		fmt.Println("=> Nothing to upload")
	} else {
		word := "files"
		if total == 1 {
			word = "file"
		}

		fmt.Printf("=> %v %s staged for upload\n", total, word)
	}

	ctx := aws.BackgroundContext()

	return s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.Concurrency = 128
		u.PartSize = UPLOAD_PART_SIZE
	}).UploadWithIterator(ctx, iter)
}
