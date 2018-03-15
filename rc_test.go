package randomchoice

import (
	"testing"
)

func Benchmark100_3(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		RandomChoice(100, 3)
	}
}

func Benchmark10_3(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		RandomChoice(10, 3)
	}
}
