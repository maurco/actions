package toolkit

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChdirFromEnv(t *testing.T) {
	d, _ := os.Getwd()
	os.Chdir("..")

	cwd, _ := os.Getwd()
	assert.NotEqual(t, cwd, d)

	os.Setenv("FOOBAR", "toolkit")
	ChdirFromEnv("FOOBAR")

	cwd, _ = os.Getwd()
	assert.Equal(t, cwd, d)
	assert.Panics(t, func() {
		os.Setenv("FOOBAR", "does-not-exist")
		ChdirFromEnv("FOOBAR")
	})
}

func TestAddFlagFromEnv(t *testing.T) {
	var flags []string
	os.Setenv("FOOBAR", "bar")

	AddFlagFromEnv(&flags, 1, "foo", "FOOBAR")
	AddFlagFromEnv(&flags, 2, "foo", "FOOBAR")
	AddFlagFromEnv(&flags, 3, "foo", "FOOBAR")
	AddFlagFromEnv(&flags, 4, "foo", "FOOBAR")

	assert.Equal(t, flags[0], "--foo=bar")
	assert.Equal(t, flags[1], "--foo bar")
	assert.Equal(t, flags[2], "-foo=bar")
	assert.Equal(t, flags[3], "-foo bar")
	assert.Panics(t, func() {
		AddFlagFromEnv(&flags, 5, "foo", "FOOBAR")
	})
}
