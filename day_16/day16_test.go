package main

import (
	"fmt"
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
	assert(doPhase(input, 0)).Equal([]int{4, 8, 2, 2, 6, 1, 5, 8})

	input = []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}
	for i := 0; i < 100; i++ {
		//fmt.Println(input)
		input = doPhase(input, 0)
	}
	fmt.Println(input)

	input = []int{3, 3, 3, 3, 3, 3, 4, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}
	for i := 0; i < 100; i++ {
		//fmt.Println(input)
		input = doPhase(input, 0)
	}
	fmt.Println(input)

	input = []int{3, 3, 3, 3, 3, 3, 4, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}
	for i := 0; i < 100; i++ {
		//fmt.Println(input)
		input = doPhase(input, 9)
	}
	fmt.Println(input)

}
