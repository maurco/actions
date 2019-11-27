package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/maurerlabs/actions/toolkit"
)

func main() {
	var (
		ignorePattern           = toolkit.GetInput("ignore_pattern", &toolkit.InputOptions{Fallback: "(.git|node_modules)"})
		bucketName              = toolkit.GetInput("bucket_name", &toolkit.InputOptions{Required: true})
		keyPrefix               = toolkit.GetInput("key_prefix")
		objectACL               = toolkit.GetInput("object_acl", &toolkit.InputOptions{Fallback: "private"})
		deleteStaleFiles        = toolkit.GetInput("delete_stale_files", &toolkit.InputOptions{Fallback: "false"}) == "true"
		cloudfrontId            = toolkit.GetInput("cloudfront_id")
		invalidateWildcard      = toolkit.GetInput("invalidate_wildcard", &toolkit.InputOptions{Fallback: "true"}) == "true"
		invalidateWithKeyPrefix = toolkit.GetInput("invalidate_with_key_prefix", &toolkit.InputOptions{Fallback: "true"}) == "true"
	)

	os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
	toolkit.ChdirFromEnv("GITHUB_WORKSPACE")
	toolkit.ChdirFromEnv("INPUT_BASE_DIR")

	sess, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}

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
}
