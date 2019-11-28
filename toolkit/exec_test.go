package toolkit

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
