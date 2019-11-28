package toolkit

import (
	"runtime"
	"testing"
)

func TestDownloadFile(t *testing.T) {

}

func TestExtractZip(t *testing.T) {

}

func TestExtractTar(t *testing.T) {

}

func TestExtract7z(t *testing.T) {

}

func TestCacheBin(t *testing.T) {

}

func TestCacheTool(t *testing.T) {
	extracted := ExtractTar(DownloadFile("https://github.com/gohugoio/hugo/releases/download/v0.60.0/hugo_extended_0.60.0_Linux-64bit.tar.gz"))
	cache := CacheTool(extracted, "hugo", "0.60.0", runtime.GOARCH)
	t.Log(cache)

	// extracted = ExtractTar(DownloadFile("https://nodejs.org/dist/v12.7.0/node-v12.7.0-linux-x64.tar.gz"))
	// cache := CacheTool(extracted, "node", "12.7.0", runtime.GOARCH)
	// t.Log(cache)
}

func TestFind(t *testing.T) {

}

func TestFindAllVersions(t *testing.T) {

}
