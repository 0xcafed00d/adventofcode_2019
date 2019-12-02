package main

import (
	"testing"

	"github.com/0xcafed00d/assert"
)

func TestExecCode(t *testing.T) {
	assert := assert.Make(t)

	assert(execCode([]int{1, 0, 0, 0, 99})).Equal([]int{2, 0, 0, 0, 99})
	assert(execCode([]int{2, 3, 0, 3, 99})).Equal([]int{2, 3, 0, 6, 99})
	assert(execCode([]int{2, 4, 4, 5, 99, 0})).Equal([]int{2, 4, 4, 5, 99, 9801})
	assert(execCode([]int{1, 1, 1, 4, 99, 5, 6, 0, 99})).Equal([]int{30, 1, 1, 4, 2, 5, 6, 0, 99})
}
