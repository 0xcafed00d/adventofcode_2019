package main

import (
	"testing"

	"github.com/0xcafed00d/assert"
)

func TestCalcFuel(t *testing.T) {
	assert := assert.Make(t)

	assert(calcFuel(12) == 2)
	assert(calcFuel(14) == 2)
	assert(calcFuel(1969) == 654)
	assert(calcFuel(100756) == 33583)
}
