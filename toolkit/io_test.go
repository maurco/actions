package toolkit

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMkdirRand(t *testing.T) {

}

func TestGoToWorkspace(t *testing.T) {
	d, _ := os.Getwd()
	os.Chdir("../../")
	cwd, _ := os.Getwd()
	assert.NotEqual(t, cwd, d)

	os.Setenv("GITHUB_WORKSPACE", "github-actions")
	os.Setenv("FOOBAR", "toolkit")
	GoToWorkspace("FOOBAR")

	cwd, _ = os.Getwd()
	assert.Equal(t, cwd, d)
	assert.Panics(t, func() {
		os.Setenv("FOOBAR", "does-not-exist")
		GoToWorkspace("FOOBAR")
	})
}
