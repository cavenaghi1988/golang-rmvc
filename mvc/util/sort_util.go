package util

import "sort"

// BabbleSort is sort number in slice
func BabbleSort(element []int) {
	keepRunning := true
	for keepRunning {
		keepRunning = false

		for i := 0; i < len(element)-1; i++ {
			if element[i] > element[i+1] {
				element[i], element[i+1] = element[i+1], element[i]
				keepRunning = true
			}
		}
	}
}

// SortElemnts change sort native and babble
func SortElemnts(el []int) {
	if len(el) < 1000 {
		BabbleSort(el)
	} else {
		sort.Ints(el)
	}

}
