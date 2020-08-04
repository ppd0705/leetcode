package main

func relativeSortArray(arr1 []int, arr2 []int) []int {
	counter := [1001]int{}
	for _, i := range arr1 {
		counter[i]++
	}

	idx := 0
	for _, i := range arr2 {
		for counter[i] > 0 {
			arr1[idx] = i
			idx++
			counter[i]--
		}
	}

	for i := range counter {
		for counter[i] > 0 {
			arr1[idx] = i
			idx++
			counter[i]--
		}
	}
	return arr1
}

func relativeSortArray2(arr1 []int, arr2 []int) []int {
	lastIdx := 0
	for i := range arr2 {
		for j := range arr1 {
			if arr2[i] == arr1[j] {
				arr1[j], arr1[lastIdx] = arr1[lastIdx], arr1[j]
				lastIdx++
			}
		}
	}
	for i := lastIdx; i < len(arr1)-1; i++ {
		for j := i + 1; j < len(arr1); j++ {
			if arr1[j] < arr1[i] {
				arr1[i], arr1[j] = arr1[j], arr1[i]
			}
		}
	}
	return arr1
}
