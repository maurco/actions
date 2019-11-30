package toolkit

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"

	"github.com/otiai10/copy"
)

const (
	IS_WINDOWS = runtime.GOOS == "windows"
	IS_DARWIN  = runtime.GOOS == "darwin"
	USER_AGENT = "maurerlabs/actions/toolkit"
)

type DownloadOptions struct {
	OutputName string
	OutputPath string
}

type ExtractOptions struct {
	KeepContainerDir  bool
	DeleteArchiveFile bool
}

func getBaseDir() string {
	if _, ok := os.LookupEnv("GITHUB_ACTIONS"); !ok {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		return filepath.Join(wd, "_")
	}
	if IS_WINDOWS {
		if val, ok := os.LookupEnv("USERPROFILE"); ok {
			return val
		}
		return "C:\\"
	}
	if IS_DARWIN {
		return "/Users"
	}
	return "/home"
}

func getTempDir() (dir string) {
	if val, ok := os.LookupEnv("RUNNER_TEMP"); ok {
		dir = val
	} else {
		dir = filepath.Join(getBaseDir(), "actions", "temp")
	}

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	return
}

func getCacheDir() (dir string) {
	if val, ok := os.LookupEnv("RUNNER_TOOL_CACHE"); ok {
		dir = val
	} else {
		dir = filepath.Join(getBaseDir(), "actions", "cache")
	}

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	return
}

func collapseContainerDir(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	if len(files) == 1 && files[0].IsDir() {
		containerDir := filepath.Join(path, files[0].Name())
		containerDirTemp := filepath.Join(filepath.Dir(path), files[0].Name()+"_TEMP")

		err = os.Rename(containerDir, containerDirTemp)
		if err != nil {
			panic(err)
		}
		err = os.Remove(path)
		if err != nil {
			panic(err)
		}
		err = os.Rename(containerDirTemp, path)
		if err != nil {
			panic(err)
		}
	}
}

func DownloadFile(url string, options ...*DownloadOptions) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", USER_AGENT)

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		panic(fmt.Sprintf("Received status code %d for %s", res.StatusCode, url))
	}

	var outPath string
	if len(options) > 0 && options[0].OutputPath != "" {
		outPath = options[0].OutputPath
	} else {
		outPath = MkdirRand(getTempDir())
	}

	var outFile string
	if len(options) > 0 && options[0].OutputName != "" {
		outFile = filepath.Join(outPath, options[0].OutputName)
	} else {
		outFile = filepath.Join(outPath, filepath.Base(url))
	}

	if _, err := os.Stat(outFile); err == nil {
		panic(fmt.Sprintf("%s already exists", outFile))
	}

	out, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		panic(err)
	}

	return outFile
}

func GunzipFile(path string, options ...*ExtractOptions) string {
	ext := filepath.Ext(path)
	if ext != ".gz" {
		panic(fmt.Sprintf("%s is not a gzipped file", path))
	}

	out := path[:len(path)-len(ext)]
	if _, err := os.Stat(out); err == nil {
		panic(fmt.Sprintf("%s already exists", out))
	}

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r, err := gzip.NewReader(f)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	o, err := os.Create(out)
	if err != nil {
		panic(err)
	}
	defer o.Close()

	for {
		r.Multistream(false)
		if _, err := io.Copy(o, r); err != nil {
			panic(err)
		}

		err := r.Reset(f)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	if len(options) > 0 && options[0].DeleteArchiveFile {
		err = os.Remove(path)
		if err != nil {
			panic(err)
		}
	}

	return out
}

func ExtractZip() {

}

func ExtractTar(path string, options ...*ExtractOptions) string {
	re := regexp.MustCompile("(.*)\\.tar(\\.gz)?$")
	match := re.FindStringSubmatch(path)
	if match == nil {
		panic(fmt.Sprintf("%s is not a tar archive", path))
	}

	out := match[1]
	if _, err := os.Stat(out); err == nil {
		panic(fmt.Sprintf("%s already exists", out))
	}

	err := os.MkdirAll(out, os.ModePerm)
	if err != nil {
		panic(err)
	}

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var r io.ReadCloser = f
	if match[2] != "" {
		if r, err = gzip.NewReader(f); err != nil {
			panic(err)
		}
		defer r.Close()
	}

	tr := tar.NewReader(r)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		} else if header == nil {
			continue
		}

		subpath := filepath.Join(out, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if _, err := os.Stat(subpath); err == nil {
				panic(fmt.Sprintf("%s already exists", subpath))
			}
			if err := os.MkdirAll(subpath, os.FileMode(header.Mode)); err != nil {
				panic(err)
			}
		case tar.TypeLink:
			if err := os.Link(header.Linkname, subpath); err != nil {
				panic(err)
			}
		case tar.TypeSymlink:
			if err := os.Symlink(header.Linkname, subpath); err != nil {
				panic(err)
			}
		case tar.TypeReg:
			fr, err := os.OpenFile(subpath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				panic(err)
			}
			if _, err := io.Copy(fr, tr); err != nil {
				panic(err)
			}
			fr.Close()
		default:
			panic(fmt.Sprintf("Unable to untar type %c in file %s", header.Typeflag, path))
		}
	}

	if options == nil || !options[0].KeepContainerDir {
		collapseContainerDir(out)
	}

	if len(options) > 0 && options[0].DeleteArchiveFile {
		err = os.Remove(path)
		if err != nil {
			panic(err)
		}
	}

	return out
}

func Extract7z() {

}

func CacheTool(path, name, version, arch string) string {
	out := filepath.Join(getCacheDir(), name, version)
	outArch := filepath.Join(out, arch)
	outComplete := outArch + ".complete"

	err := os.RemoveAll(out)
	if err != nil {
		panic(err)
	}
	err = os.RemoveAll(outComplete)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(out, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fi, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	if fi.IsDir() {
		err = copy.Copy(path, outArch)
		if err != nil {
			panic(err)
		}
	} else {
		err = copy.Copy(path, filepath.Join(outArch, filepath.Base(path)))
		if err != nil {
			panic(err)
		}
	}

	err = ioutil.WriteFile(outComplete, nil, os.ModePerm)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s%c", outArch, os.PathSeparator)
}

func Find() {

}

func FindAllVersions() {

}

func InstallBin(paths ...string) {
	bin := "bin"
	path := filepath.Join(paths...)
	name := filepath.Base(path)

	info, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	home := os.Getenv("HOME")
	if home == "" {
		panic("Cannot find home directory")
	}

	pathBin := filepath.Join(home, bin)
	if err := os.MkdirAll(pathBin, os.ModePerm); err != nil {
		panic(err)
	}

	pathNew := filepath.Join(pathBin, name)
	if err := copy.Copy(path, pathNew); err != nil {
		panic(err)
	}

	if !info.IsDir() {
		if err := os.Chmod(pathNew, 0777); err != nil {
			panic(err)
		}
	}

	AddPath(pathBin)
	AddPath(fmt.Sprintf("/home/runner/work/_temp/_github_home/%s", bin))
}
