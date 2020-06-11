package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadFile() []string {
	dat, err := ioutil.ReadFile("./input1k.txt")
	check(err)

	return strings.Split(string(dat), " ")
}

func BenchmarkEchoWithJoin(b *testing.B) {
	dat := loadFile()

	var s, sep string

	for i := 1; i < len(dat); i++ {
		s += sep + dat[i]
		sep = " "
	}
}

func BenchmarkEchoWithFor(b *testing.B) {
	dat := loadFile()

	strings.Join(dat, " ")
}
