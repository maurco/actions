package actions

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChangeDirByEnvVar(t *testing.T) {
	d, _ := os.Getwd()
	os.Chdir("..")

	cwd, _ := os.Getwd()
	assert.NotEqual(t, cwd, d)

	os.Setenv("FOOBAR", "actions")
	ChangeDirByEnvVar("FOOBAR")

	cwd, _ = os.Getwd()
	assert.Equal(t, cwd, d)
	assert.Panics(t, func() {
		os.Setenv("FOOBAR", "does-not-exist")
		ChangeDirByEnvVar("FOOBAR")
	})
}

func TestAddFlagByEnvVar(t *testing.T) {
	var flags []string
	os.Setenv("FOOBAR", "bar")

	AddFlagByEnvVar(&flags, 1, "foo", "FOOBAR")
	AddFlagByEnvVar(&flags, 2, "foo", "FOOBAR")
	AddFlagByEnvVar(&flags, 3, "foo", "FOOBAR")
	AddFlagByEnvVar(&flags, 4, "foo", "FOOBAR")

	assert.Equal(t, flags[0], "--foo=bar")
	assert.Equal(t, flags[1], "--foo bar")
	assert.Equal(t, flags[2], "-foo=bar")
	assert.Equal(t, flags[3], "-foo bar")
	assert.Panics(t, func() {
		AddFlagByEnvVar(&flags, 5, "foo", "FOOBAR")
	})
}
