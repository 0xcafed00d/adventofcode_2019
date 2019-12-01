package main

import (
	"testing"

	"github.com/0xcafed00d/assert"
)

func TestCalcFuel(t *testing.T) {
	assert := assert.Make(t)

	assert(calcFuel(12)).Equal(2)
	assert(calcFuel(14)).Equal(2)
	assert(calcFuel(1969)).Equal(654)
	assert(calcFuel(100756)).Equal(33583)
}
