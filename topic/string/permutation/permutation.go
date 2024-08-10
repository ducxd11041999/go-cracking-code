package permutation

import "sort"

// Time complexcity dependence for sort algorithm
func IsPermutationWithSort(str1, str2 string) bool {
	s1Sorted := []rune(str1)
	sort.Slice(s1Sorted, func(i, j int) bool { //sort the string using the function
		return s1Sorted[i] < s1Sorted[j]
	})

	s2Sorted := []rune(str2)
	sort.Slice(s2Sorted, func(i, j int) bool { //sort the string using the function
		return s2Sorted[i] < s2Sorted[j]
	})

	return string(s1Sorted) == string(s2Sorted)
}

// O(n)
func IsPermutationWithCount(str1, str2 string) bool {
	counter := make([]int, 256)
	le1 := len(str1)
	le2 := len(str2)

	if le1 != le2 {
		return false
	}

	for i := 0; i < le1; i++ {
		counter[str1[i]]++
	}

	for i := 0; i < le2; i++ {
		exist := counter[str1[i]] - 1
		if exist < 0 {
			return false
		}
	}

	return true
}
