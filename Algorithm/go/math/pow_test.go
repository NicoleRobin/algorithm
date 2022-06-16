package math

import (
	"math"
	"testing"
)

func TestPow(t *testing.T) {
	result := math.Pow(1, 5)
	if result != 1 {
		t.Fatalf("want:%f, get:%f\n", 1.0, result)
	}
}
