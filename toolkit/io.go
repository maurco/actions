package toolkit

import (
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func MakeExecutable(paths ...string) {
	path := filepath.Join(paths...)

	if err := os.Chmod(path, 0755); err != nil {
		panic(err)
	}
}

func MkdirRand(base string) string {
	length := 10
	chars := []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	rand.Seed(time.Now().UnixNano())

	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	dir := filepath.Join(base, b.String())
	if _, err := os.Stat(dir); err == nil {
		return MkdirRand(base)
	}

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	return dir
}

func GoToWorkspace(keys ...string) {
	if val, ok := os.LookupEnv("GITHUB_WORKSPACE"); ok && val != "" {
		err := os.Chdir(val)
		if err != nil {
			panic(err)
		}
	}

	for _, v := range keys {
		if val, ok := os.LookupEnv(v); ok && val != "" {
			err := os.Chdir(val)
			if err != nil {
				panic(err)
			}
		}
	}
}
