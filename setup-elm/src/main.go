package main

import (
	"fmt"

	"github.com/maurerlabs/actions/toolkit"
)

const LATEST_VERSION = "0.19.1"

func main() {
	var (
		version = toolkit.GetInput("version", &toolkit.InputOptions{Fallback: LATEST_VERSION})
	)

	toolkit.Info("Downloading Elm")

	url := fmt.Sprintf("https://github.com/elm/compiler/releases/download/%s/binary-for-linux-64-bit.gz", version)
	file := toolkit.DownloadFile(url, &toolkit.DownloadOptions{OutputName: "elm.gz"})

	extracted := toolkit.GunzipFile(file)
	cache := toolkit.CacheTool(extracted, "elm", version, "amd64")

	toolkit.InstallBin(cache, "elm")
	toolkit.Command("elm", "--version")
}
