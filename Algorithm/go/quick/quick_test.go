package quick

import "testing"

func TestQuickMul(t *testing.T) {
	x, y := 5, 53
	want := 265
	ans := QuickMul(x, y)

	if ans != want {
		t.Fatalf("got:%d, want:%d\n", ans, want)
	}
}

func TestQuickPow(t *testing.T) {
	x, y := 2, 10
	want := 1024
	ans := QuickPow(x, y)

	if ans != want {
		t.Fatalf("got:%d, want:%d\n", ans, want)
	}
}
