package main

import "testing"

func TestAdd(t *testing.T) {
	sumInt := Add(3, 4)
	wantInt := 7

	if sumInt == wantInt {
		t.Logf("Add(3, 4) = %v, PASSED, expected %v, got %v", sumInt, wantInt, sumInt)
	} else {
		t.Fatalf(`Add(3, 4) = %v, want match for %v`, sumInt, wantInt)
	}
}
