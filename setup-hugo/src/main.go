package main

import (
	"fmt"

	"github.com/maurerlabs/actions/toolkit"
)

const LATEST_VERSION = "0.60.0"

func main() {
	var (
		version  = toolkit.GetInput("version", &toolkit.InputOptions{Fallback: LATEST_VERSION})
		extended = toolkit.GetInput("extended", &toolkit.InputOptions{Fallback: "true"}) == "true"
	)

	var extForURL string
	if extended {
		extForURL = "_extended"
	}

	var extForLog string
	if extended {
		extForLog = " Extended"
	}

	toolkit.Info("Downloading Hugo%s", extForLog)

	url := fmt.Sprintf("https://github.com/gohugoio/hugo/releases/download/v%s/hugo%s_%s_Linux-64bit.tar.gz", version, extForURL, version)
	name := fmt.Sprintf("hugo%s", extForURL)

	extracted := toolkit.ExtractTar(toolkit.DownloadFile(url))
	cache := toolkit.CacheTool(extracted, name, version, "amd64")

	toolkit.AddPath(cache)
	toolkit.MakeExecutable(cache, "hugo")

	toolkit.Info("Installed Hugo%s v%s", extForLog, version)
}
