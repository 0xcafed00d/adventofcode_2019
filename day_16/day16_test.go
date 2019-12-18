package main

import (
	"testing"

	"github.com/0xcafed00d/assert"
)

func Test1(t *testing.T) {
	assert := assert.Make(t)

	pgen := pattern(0)
	assert(pgen()).Equal(1)
	assert(pgen()).Equal(0)
	assert(pgen()).Equal(-1)
	assert(pgen()).Equal(0)

	pgen = pattern(1)
	assert(pgen()).Equal(0)
	assert(pgen()).Equal(1)
	assert(pgen()).Equal(1)
	assert(pgen()).Equal(0)
	assert(pgen()).Equal(0)
	assert(pgen()).Equal(-1)
	assert(pgen()).Equal(-1)
	assert(pgen()).Equal(0)
	assert(pgen()).Equal(0)

	pgen = pattern(2)
	assert(pgen()).Equal(0)
	assert(pgen()).Equal(0)
	assert(pgen()).Equal(1)
	assert(pgen()).Equal(1)
	assert(pgen()).Equal(1)
	assert(pgen()).Equal(0)
	assert(pgen()).Equal(0)
	assert(pgen()).Equal(0)
	assert(pgen()).Equal(-1)
	assert(pgen()).Equal(-1)
	assert(pgen()).Equal(-1)
	assert(pgen()).Equal(0)
	assert(pgen()).Equal(0)
	assert(pgen()).Equal(0)
}

func Test2(t *testing.T) {
	assert := assert.Make(t)

	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	assert(doPhase(input)).Equal([]int{4, 8, 2, 2, 6, 1, 5, 8})
}
