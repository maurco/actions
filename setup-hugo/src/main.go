package main

import (
	"fmt"

	"github.com/maurerlabs/actions/toolkit"
)

const LATEST_VERSION = "0.60.0"

func main() {
	var (
		version  = toolkit.GetInput("version", &toolkit.InputOptions{Fallback: LATEST_VERSION})
		extended = toolkit.GetInput("extended", &toolkit.InputOptions{Fallback: "false"}) == "true"
	)

	if extended {
		toolkit.Info("Downloading Hugo Extended")
	} else {
		toolkit.Info("Downloading Hugo")
	}

	name := "hugo"
	if extended {
		name = "hugo_extended"
	}

	url := fmt.Sprintf("https://github.com/gohugoio/hugo/releases/download/v%s/%s_%s_Linux-64bit.tar.gz", version, name, version)
	extracted := toolkit.ExtractTar(toolkit.DownloadFile(url))
	cache := toolkit.CacheTool(extracted, name, version, "amd64")

	toolkit.MakeExecutable(cache, "hugo")
	toolkit.CacheBin(cache, "hugo")

	toolkit.Command("hugo", "version")
}
