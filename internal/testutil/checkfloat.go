package testutil

import "testing"

func CheckFloat(t *testing.T, name string, got, want, tol float64) {
	t.Helper()
	diff := got - want
	if diff < 0 {
		diff = -diff
	}
	if diff > tol {
		t.Fatalf("%s = %.10f, want %.10f", name, got, want)
	}
}
