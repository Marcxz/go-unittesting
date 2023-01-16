package sort

import (
	"sort"
	_ "sort"
)

func BubbleSort(elements []int) {
	keepWorking := true
	for keepWorking {
		keepWorking = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] < elements[i+1] {
				keepWorking = true
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
	}
}

func Sort(elements []int) {
	sort.Ints(elements)
}

func Greeter() string {
	return "OK"
}
