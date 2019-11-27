package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/maurerlabs/actions/toolkit"
)

func main() {
	var (
		start                   = time.Now()
		ignorePattern           = os.Getenv("INPUT_IGNORE_PATTERN")
		bucketName              = os.Getenv("INPUT_BUCKET_NAME")
		keyPrefix               = os.Getenv("INPUT_KEY_PREFIX")
		objectACL               = os.Getenv("INPUT_OBJECT_ACL")
		deleteStaleFiles        = os.Getenv("INPUT_DELETE_STALE_FILES") == "true"
		cloudfrontId            = os.Getenv("INPUT_CLOUDFRONT_ID")
		invalidateWildcard      = os.Getenv("INPUT_INVALIDATE_WILDCARD") == "true"
		invalidateWithKeyPrefix = os.Getenv("INPUT_INVALIDATE_WITH_KEY_PREFIX") == "true"
	)

	os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
	toolkit.ChdirFromEnv("GITHUB_WORKSPACE")
	toolkit.ChdirFromEnv("INPUT_BASE_DIR")

	sess, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=> Generating file lists")

	localFiles, err := getLocalFiles(".", ignorePattern)
	if err != nil {
		log.Fatal(err)
	}

	remoteFiles, err := getRemoteFiles(sess, bucketName, keyPrefix)
	if err != nil {
		log.Fatal(err)
	}

	stagedUploads, stagedDeletes := getStaged(localFiles, remoteFiles)

	err = batchUpload(sess, stagedUploads, bucketName, keyPrefix, objectACL)
	if err != nil {
		log.Fatal(err)
	}

	if deleteStaleFiles {
		err = batchDelete(sess, stagedDeletes, bucketName, keyPrefix)
		if err != nil {
			log.Fatal(err)
		}
	}

	if cloudfrontId != "" {
		var prefix string
		if invalidateWithKeyPrefix {
			prefix = keyPrefix
		}

		paths := append(stagedUploads, stagedDeletes...)
		err := invalidate(sess, cloudfrontId, invalidateWildcard, prefix, paths)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("=> Done in %s\n", time.Since(start))
}
