package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/dustin/go-humanize"
	"github.com/maurerlabs/actions/toolkit"
)

func generatePaths(wildcard bool, prefix string, paths []*string) ([]*string, int64) {
	var items []*string
	var length int

	// add leading slash
	if prefix == "" {
		prefix = "/"
	} else if prefix[0] != '/' {
		prefix = "/" + prefix
	}

	if wildcard && len(paths) > 0 {
		itemWild := prefix + "*"
		items = []*string{&itemWild}
		length = 1
	} else {
		for _, v := range paths {
			path := prefix + *v
			items = append(items, &path)

			indexHtml := regexp.MustCompile("(.*)\\/index\\.html$")
			if indexHtml.MatchString(path) {
				r := indexHtml.ReplaceAllString(path, "/$1")
				items = append(items, &r)
			}
		}

		length = len(items)
	}

	return items, int64(length)
}

func invalidate(sess *session.Session, id string, wildcard bool, prefix string, paths []*string) error {
	items, length := generatePaths(wildcard, prefix, paths)

	if length < 1 {
		toolkit.Info("Nothing to invalidate")
		return nil
	}

	word := "paths"
	if length == 1 {
		word = "path"
	}

	toolkit.StartGroup(fmt.Sprintf("%d %s staged for invalidation", length, word))
	defer toolkit.EndGroup()

	if !wildcard && length > 50 {
		cost := float64(length) * 0.005
		toolkit.Warning(fmt.Sprintf("Invalidation will cost ~$%v (%d paths), consider using a wildcard", humanize.CommafWithDigits(cost, 2), length))
	}

	ref := strconv.FormatInt(time.Now().UnixNano(), 10)
	res, err := cloudfront.New(sess).CreateInvalidation(&cloudfront.CreateInvalidationInput{
		DistributionId: &id,
		InvalidationBatch: &cloudfront.InvalidationBatch{
			CallerReference: &ref,
			Paths: &cloudfront.Paths{
				Quantity: &length,
				Items:    items,
			},
		},
	})
	if err != nil {
		return err
	}

	for _, v := range res.Invalidation.InvalidationBatch.Paths.Items {
		toolkit.Info(fmt.Sprintf("Invalidated %s", *v))
	}

	return nil
}
