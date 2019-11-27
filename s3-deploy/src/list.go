package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type FileMap = map[string]string

var fileMapLock = sync.Mutex{}
var fileMapWG = sync.WaitGroup{}

func generateEtag(path string, files *FileMap) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var chunks [][]byte
	for {
		buf := make([]byte, UPLOAD_PART_SIZE)

		if _, err := f.Read(buf); err != nil {
			if err != io.EOF {
				panic(err)
			}
		}

		// If file or chunk is smaller than the part size,
		// make sure to trim any empty space in the buffer
		// since it will affect the hash.
		buf = bytes.Trim(buf, "\x00")

		hash := md5.New()
		hash.Write(buf)
		chunks = append(chunks, hash.Sum(nil))

		if err == io.EOF || len(buf) < UPLOAD_PART_SIZE {
			break
		}
	}

	sum := hex.EncodeToString(chunks[0])
	if len(chunks) > 1 {
		concat := bytes.Join(chunks, []byte(""))
		sum = fmt.Sprintf("%x-%d", md5.Sum(concat), len(chunks))
	}

	fileMapLock.Lock()
	defer fileMapLock.Unlock()
	defer fileMapWG.Done()

	(*files)[path] = sum
}

func getLocalFiles(dir string, ignore interface{}) (*FileMap, error) {
	files := make(FileMap)

	// Prevent regexp from being re-compiled for every subdirectory
	switch v := ignore.(type) {
	case *regexp.Regexp:
		// Do nothing
	case string:
		if ignore != "" {
			ignore = regexp.MustCompile(v)
		}
	default:
		return nil, errors.New(fmt.Sprintf("Expected a string for ignore pattern, got %T", v))
	}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		pattern, ok := ignore.(*regexp.Regexp)
		if ok && pattern.MatchString(path) {
			return nil
		}

		if info.IsDir() && path != "." {
			getLocalFiles(filepath.Join(dir, path), ignore)
		} else if path != "." {
			fileMapWG.Add(1)
			go generateEtag(path, &files)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	fileMapWG.Wait()
	return &files, nil
}

func getRemoteFiles(sess *session.Session, bucket, prefix string) (*FileMap, error) {
	files := make(FileMap)

	err := s3.New(sess).ListObjectsV2Pages(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix),
	}, func(p *s3.ListObjectsV2Output, last bool) bool {
		for _, obj := range p.Contents {
			key, etag := *obj.Key, *obj.ETag

			// Remove surrounding quotes
			if len(etag) > 0 && etag[0] == '"' {
				etag = etag[1:]
			}
			if len(etag) > 0 && etag[len(etag)-1] == '"' {
				etag = etag[:len(etag)-1]
			}

			// Remove prefix
			if prefix != "" {
				if prefix[len(prefix)-1] == '/' {
					key = key[len(prefix):]
				} else {
					key = key[len(prefix)+1:]
				}
			}

			files[key] = etag
		}

		return true
	})

	return &files, err
}
