package main

import (
	"testing"

	"github.com/0xcafed00d/assert"
)

func Test2(t *testing.T) {
	assert := assert.Make(t)

	cpu := intcodecpu{}
	cpu.setProgram([]int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99})
	cpu.execCode()
	assert(cpu.output).Equal([]int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99})

	cpu = intcodecpu{}
	cpu.setProgram([]int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0})
	cpu.execCode()
	assert(cpu.output).Equal([]int64{34915192 * 34915192})

	cpu = intcodecpu{}
	cpu.setProgram([]int64{104, 1125899906842624, 99})
	cpu.execCode()
	assert(cpu.output).Equal([]int64{1125899906842624})
}
