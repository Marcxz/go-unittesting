package services

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	elements := []int{3, 8, 7, 6, 10}
	Sort(elements)

	fmt.Println(elements)

	if elements[0] != 10 {
		t.Error("First element should be 10")
	}
	if elements[len(elements)-1] != 3 {
		t.Error("Last element should be 3")
	}
}
