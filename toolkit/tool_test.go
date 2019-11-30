package toolkit

import (
	// "runtime"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	// file := DownloadFile("https://github.com/gohugoio/hugo/releases/download/v0.60.0/hugo_extended_0.60.0_Linux-64bit.tar.gz")
	// t.Log(file)
}

func TestGunzipFile(t *testing.T) {
	// file := DownloadFile("https://github.com/elm/compiler/releases/download/0.19.1/binary-for-linux-64-bit.gz", &DownloadOptions{OutputName: "elm.gz"})
	// extracted := GunzipFile(file)
	// t.Log(extracted)
}

func TestExtractZip(t *testing.T) {

}

func TestExtractTar(t *testing.T) {
	// file := DownloadFile("https://nodejs.org/dist/v12.7.0/node-v12.7.0-linux-x64.tar.gz")
	// extracted := ExtractTar(file)
	// t.Log(extracted)
}

func TestExtract7z(t *testing.T) {

}

func TestCacheTool(t *testing.T) {
	// extracted := ExtractTar()
	// cache := CacheTool(extracted, "hugo", "0.60.0", runtime.GOARCH)
	// t.Log(cache)

	// file := DownloadFile("https://github.com/elm/compiler/releases/download/0.19.1/binary-for-linux-64-bit.gz", &DownloadOptions{OutputName: "elm.gz"})
	// extracted := GunzipFile(file)
	// cache := CacheTool(extracted, "elm", "0.19.1", runtime.GOARCH)
	// t.Log(cache)
}

func TestFind(t *testing.T) {

}

func TestFindAllVersions(t *testing.T) {

}

func TestInstallBin(t *testing.T) {

}
