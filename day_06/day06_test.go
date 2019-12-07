package main

import (
	"testing"

	"github.com/0xcafed00d/assert"
)

func TestExecCode(t *testing.T) {
	assert := assert.Make(t)

	s := newSystem()
	s.add("AAA", "COM")
	s.add("BBB", "AAA")
	assert(s.getChainLen("BBB")).Equal(2)

	s = newSystem()
	s.add("B", "COM")
	s.add("C", "B")
	s.add("D", "C")
	s.add("E", "D")
	s.add("F", "E")
	s.add("G", "B")
	s.add("H", "G")
	s.add("I", "D")
	s.add("J", "E")
	s.add("K", "J")
	s.add("L", "K")
	assert(s.getOrbitCount()).Equal(42)

	s.add("YOU", "K")
	s.add("SAN", "I")
	assert(s.numTransfers("YOU", "SAN")).Equal(4)
}
