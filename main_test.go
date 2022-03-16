package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGoFiles(t *testing.T) {
	goFiles, err := getGoFiles("test_dir_no_go_files")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(goFiles))

	goFiles, err = getGoFiles("test_dir")
	assert.NoError(t, err)
	assert.Equal(t, 4, len(goFiles))
}
