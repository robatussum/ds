package test

import (
	"testing"

	"github.com/robatussum/go_ds/alg"
)

func TestBUMergesort(t *testing.T) {
	var cases = []struct {
		i []int
	}{
		{},
		{},
		{},
	}

	for _, c := range cases {
		alg.BottomUpMerge(c.i)

		if !isSorted(c.i) {
			t.Error("E: Not sorted %v", c.i)
		}
	}
}

func isSorted(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}