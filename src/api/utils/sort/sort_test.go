package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortNoError(t *testing.T) {
	// Init
	elements := []int{1, 7, 5, 3, 1}
	expected := []int{7, 5, 3, 1, 1}
	// Execution
	BubbleSort(elements)
	assert.EqualValues(t, 7, elements[0])
	// Validation
	if elements[0] != 7 {
		t.Errorf("Failed Bubble sort expected %v, got %v", expected, elements)
	} else {
		t.Logf("Success Bubble sort expected %v, got %v", expected, elements)
	}
}

func TestSortNoError(t *testing.T) {
	// Init
	elements := []int{1, 7, 5, 3, 1}
	expected := []int{7, 5, 3, 1, 1}
	// Execution
	Sort(elements)

	// Validation
	if elements[0] != 7 {
		t.Errorf("Failed Bubble sort expected %v, got %v", expected, elements)
	} else {
		t.Logf("Success Bubble sort expected %v, got %v", expected, elements)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := []int{1, 7, 5, 3, 1}
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := []int{1, 7, 5, 3, 1}
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
