package toolkit

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeExecutable(t *testing.T) {
	os.MkdirAll("_", os.ModePerm)
	ioutil.WriteFile("_/test.bash", []byte("#!/bin/bash\n"), 0666)
	defer os.RemoveAll("_/test.bash")

	cmd := exec.Command("./_/test.bash")
	err := cmd.Run()
	assert.Contains(t, err.Error(), "denied")

	MakeExecutable("_", "test.bash")
	err = cmd.Run()
	assert.Nil(t, err)
}

func TestMkdirRand(t *testing.T) {

}

func TestGoToWorkspace(t *testing.T) {
	d, _ := os.Getwd()
	os.Chdir("../../")
	cwd, _ := os.Getwd()
	assert.NotEqual(t, cwd, d)

	os.Setenv("GITHUB_WORKSPACE", "actions")
	GoToWorkspace("toolkit")

	cwd, _ = os.Getwd()
	assert.Equal(t, cwd, d)
	assert.Panics(t, func() {
		GoToWorkspace("does-not-exist")
	})
}
