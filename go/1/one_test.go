package one

import (
	"testing"
)

func TestDistanceSum(t *testing.T) {
	sum := distanceSum([]int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3})

	if sum != 11 {
		t.Errorf("Expected 11, got %d", sum)
	}
}
