package main

import (
	"testing"

	"github.com/0xcafed00d/assert"
)

func TestCheckPW(t *testing.T) {
	assert := assert.Make(t)

	assert(getDigit(123456, 0)).Equal(6)
	assert(getDigit(123456, 1)).Equal(5)
	assert(getDigit(123456, 2)).Equal(4)
	assert(getDigit(123456, 3)).Equal(3)
	assert(getDigit(123456, 4)).Equal(2)
	assert(getDigit(123456, 5)).Equal(1)

	assert(checkPW_part1(111111)).Equal(true)
	assert(checkPW_part1(223450)).Equal(false)
	assert(checkPW_part1(123789)).Equal(false)

	assert(checkPW_part2(112233)).Equal(true)
	assert(checkPW_part2(123444)).Equal(false)
	assert(checkPW_part2(111122)).Equal(true)
}
