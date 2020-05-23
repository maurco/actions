package main

import (
	"fmt"
	"mime"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/dustin/go-humanize"
	"github.com/maurco/actions/toolkit"
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
	ext := filepath.Ext(iter.next.path)
	mimeType := mime.TypeByExtension(ext)
	key := iter.prefix + iter.next.path

	return s3manager.BatchUploadObject{
		Object: &s3manager.UploadInput{
			ACL:         &iter.acl,
			Bucket:      &iter.bucket,
			Key:         &key,
			ContentType: &mimeType,
			Body:        iter.next.f,
		},
		After: func() error {
			fi, err := iter.next.f.Stat()
			if err != nil {
				return err
			}

			size := humanize.Bytes(uint64(fi.Size()))
			toolkit.Info("Uploaded %s [%v]", key, size)

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
	title := "Nothing to upload"

	if total > 0 {
		word := "files"
		if total == 1 {
			word = "file"
		}
		title = fmt.Sprintf("Uploading %v %s", total, word)
	}

	toolkit.StartGroup(title)
	defer toolkit.EndGroup()

	if total < 1 {
		return nil
	}

	ctx := aws.BackgroundContext()
	return s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.Concurrency = 128
		u.PartSize = UPLOAD_PART_SIZE
	}).UploadWithIterator(ctx, iter)
}
