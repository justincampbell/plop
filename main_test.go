package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isTTY(t *testing.T) {
	assert.True(t, isTTY(os.Stdin))

	file, _ := ioutil.TempFile(os.TempDir(), "test")
	assert.False(t, isTTY(file))
}
