package main

import (
	"fmt"
	"strings"
	"testing"
)

var (
	testData = []string{"taylor", "rock"}
)

// run this command to check benchmark :
// go test -benchmem -bench=.

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := strings.Join(testData, ":")
		_ = s
	}
}

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s:%s", testData[0], testData[1])
		_ = s
	}
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := testData[0] + ":" + testData[1]
		_ = s
	}
}
