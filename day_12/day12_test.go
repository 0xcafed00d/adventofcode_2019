package main

import (
	"testing"

	"github.com/0xcafed00d/assert"
)

func Test2(t *testing.T) {
	assert := assert.Make(t)

	input := []point{
		point{x: -1, y: 0, z: 2},
		point{x: 2, y: -10, z: -7},
		point{x: 4, y: -8, z: 8},
		point{x: 3, y: 5, z: -1},
	}

	calcVel(input)
	applyVel(input)

	assert(input[0]).Equal(point{x: 2, y: -1, z: 1, dx: 3, dy: -1, dz: -1})
	assert(input[1]).Equal(point{x: 3, y: -7, z: -4, dx: 1, dy: 3, dz: 3})
	assert(input[2]).Equal(point{x: 1, y: -7, z: 5, dx: -3, dy: 1, dz: -3})
	assert(input[3]).Equal(point{x: 2, y: 2, z: 0, dx: -1, dy: -3, dz: 1})

	calcVel(input)
	applyVel(input)

	assert(input[0]).Equal(point{x: 5, y: -3, z: -1, dx: 3, dy: -2, dz: -2})
	assert(input[1]).Equal(point{x: 1, y: -2, z: 2, dx: -2, dy: 5, dz: 6})
	assert(input[2]).Equal(point{x: 1, y: -4, z: -1, dx: 0, dy: 3, dz: -6})
	assert(input[3]).Equal(point{x: 1, y: -4, z: 2, dx: -1, dy: -6, dz: 2})

}
