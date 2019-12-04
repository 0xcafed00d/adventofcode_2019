package main

import (
	"testing"

	"github.com/0xcafed00d/assert"
)

func TestDist(t *testing.T) {
	assert := assert.Make(t)

	assert(calcDist([]string{"R8,U5,L5,D3",
		"U7,R6,D4,L4"})).Equal(6, 30)

	assert(calcDist([]string{"R75,D30,R83,U83,L12,D49,R71,U7,L72",
		"U62,R66,U55,R34,D71,R55,D58,R83"})).Equal(159, 610)

	assert(calcDist([]string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
		"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"})).Equal(135, 410)
}
