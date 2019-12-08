package main

import (
	"testing"

	"github.com/0xcafed00d/assert"
)

func TestExecCode(t *testing.T) {
	assert := assert.Make(t)

	cpu := intcodecpu{}
	cpu.memory = []int{1, 0, 0, 0, 99}
	cpu.execCode()
	assert(cpu.memory).Equal([]int{2, 0, 0, 0, 99})

	cpu = intcodecpu{}
	cpu.memory = []int{2, 3, 0, 3, 99}
	cpu.execCode()
	assert(cpu.memory).Equal([]int{2, 3, 0, 6, 99})

	cpu = intcodecpu{}
	cpu.memory = []int{2, 4, 4, 5, 99, 0}
	cpu.execCode()
	assert(cpu.memory).Equal([]int{2, 4, 4, 5, 99, 9801})

	cpu = intcodecpu{}
	cpu.memory = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	cpu.execCode()
	assert(cpu.memory).Equal([]int{30, 1, 1, 4, 2, 5, 6, 0, 99})
}

func TestAmplifiers(t *testing.T) {
	assert := assert.Make(t)

	val := calcAmplitude([]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}, []int{4, 3, 2, 1, 0})
	assert(val).Equal(43210)

	val = calcAmplitude([]int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23,
		101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}, []int{0, 1, 2, 3, 4})
	assert(val).Equal(54321)

	val = calcAmplitude([]int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
		1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}, []int{1, 0, 4, 3, 2})
	assert(val).Equal(65210)
}

func TestAmplifiers2(t *testing.T) {
	assert := assert.Make(t)

	val := calcAmplitudeFeedback([]int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26,
		27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}, []int{9, 8, 7, 6, 5})
	assert(val).Equal(139629729)

	val = calcAmplitudeFeedback([]int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54,
		-5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4,
		53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10}, []int{9, 7, 8, 5, 6})
	assert(val).Equal(18216)
}
