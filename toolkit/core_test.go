package toolkit

import (
	"bytes"
	"io"
	"log"
	"os"
	"regexp"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: PR this into testify
func captureOutput(f func()) string {
	wg := sync.WaitGroup{}
	out := make(chan string)

	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()

	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)

	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	wg.Wait()
	f()
	writer.Close()

	return <-out
}

func TestExportVar(t *testing.T) {
	output := captureOutput(func() {
		ExportVar("foo ", "bar")
		assert.Equal(t, os.Getenv("foo"), "bar")
	})
	assert.Equal(t, output, "::set-env name=foo::bar\n")
}

func TestAddPath(t *testing.T) {
	output := captureOutput(func() {
		AddPath("/foo/bar")
		assert.Regexp(t, regexp.MustCompile("^\\/foo\\/bar:"), os.Getenv("PATH"))
	})
	assert.Equal(t, output, "::add-path::/foo/bar\n")
}

func TestGetInput(t *testing.T) {
	os.Setenv("INPUT_FOO_BAR", "baz")
	assert.Equal(t, GetInput("foo_bar"), "baz")
	assert.Equal(t, GetInput("foo bar"), "baz")
	assert.Equal(t, GetInput(" foo_bar "), "baz")
	assert.Equal(t, GetInput("FOO_BAR"), "baz")
	assert.Equal(t, GetInput("foo"), "")
	assert.Equal(t, GetInput("foo", &InputOptions{Fallback: "bar"}), "bar")
	assert.Panics(t, func() {
		GetInput("foo", &InputOptions{Required: true})
	})
}

func TestGetState(t *testing.T) {
	os.Setenv("STATE_foo_bar", "baz")
	assert.Equal(t, GetState("foo_bar"), "baz")
	assert.Equal(t, GetState("FOO_BAR"), "")
	assert.Equal(t, GetState("foo", &InputOptions{Fallback: "bar"}), "bar")
	assert.Panics(t, func() {
		GetState("foo", &InputOptions{Required: true})
	})
}
