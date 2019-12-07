package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func panicOnErr(err error, msg string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, msg)
		panic(err)
	}
}

type body struct {
	name     string
	orbiting string
}

type system map[string]body

func newSystem() system {
	s := make(map[string]body)
	s["COM"] = body{"COM", ""}
	return s
}

func (s *system) add(bodyName string, orbiting string) {
	(*s)[bodyName] = body{bodyName, orbiting}
}

func (s *system) getChainLen(bodyName string) int {
	l := 0
	for {
		o := (*s)[bodyName].orbiting
		if o == "" {
			return l
		}
		l++
		bodyName = o
	}
}

func (s *system) getOrbitCount() int {
	count := 0
	for k := range *s {
		count += (*s).getChainLen(k)
	}
	return count
}

func getIndex(val string, slice []string) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}

func (s *system) numTransfers(from string, to string) int {
	pathToCOM := []string{}
	for {
		o := (*s)[to].orbiting
		if o == "" {
			break
		}
		fmt.Println(o)
		pathToCOM = append(pathToCOM, o)
		to = o
	}

	count := 0
	for {
		o := (*s)[from].orbiting
		if o == "" {
			break
		}
		intersect := getIndex(o, pathToCOM)
		if intersect != -1 {
			return count + len(pathToCOM[0:intersect])
		}
		count++
		from = o
	}

	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	s := newSystem()

	for scanner.Scan() {
		line := scanner.Text()
		bodies := strings.Split(line, ")")
		s.add(bodies[1], bodies[0])
	}

	fmt.Printf("total orbit count: %v \n", s.getOrbitCount())
	fmt.Printf("transfer count: %v \n", s.numTransfers("YOU", "SAN"))
}
