package module02

import "sort"

// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
//
// Big O: O(N^2), where N is the size of the list
func BubbleSortInt(list []int) {
	for i := 0; i < len(list); i++ {
		var swap = false
		for j := 0; j < len(list) - 1 - i; j++ {
			if list[j] > list[j + 1] {
				list[j], list[j+1] = list[j+1], list[j]
				swap = true
			}
		}
		if !swap {
			return
		}
	}
}

// BubbleSortString is a bubble sort for string slices. Try implementing it for
// practice.
func BubbleSortString(list []string) {
	
}

// BubbleSortPerson uses bubble sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func BubbleSortPerson(people []Person) {
}

// BubbleSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice.
func BubbleSort(list sort.Interface) {
	for i := 0; i < list.Len(); i++ {
		var swap = false
		for j := 0; j < list.Len() - 1 - i; j++ {
			if !list.Less(j, j+1) && !list.Less(j+1, j) {
				// do nothing
			} else if !list.Less(j, j+1) {
				list.Swap(j, j+1)
				swap = true
			}
		}
		if !swap {
			return
		}
	}
}
