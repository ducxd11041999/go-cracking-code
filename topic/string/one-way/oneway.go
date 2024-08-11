package oneway

func OneWay(str1, str2 string) bool {
	le1 := len(str1)
	le2 := len(str2)

	if le1 == le2 {
		return oneEditReplace(str1, str2, le1, le2)
	}

	if le1 == le2-1 {
		return oneEdit(str1, str2, le1, le2)
	}

	return oneEdit(str2, str1, le2, le1)
}

func OneWayCommon(str1, str2 string) bool {
	le1 := len(str1)
	le2 := len(str2)

	if abs(le1-le2) > 1 {
		return false
	}

	if le1 < le2 {
		return merge2Case(str1, str2)
	}

	return merge2Case(str2, str1)
}

func oneEditReplace(str1, str2 string, le1, le2 int) bool {
	foundDiff := false
	idx1, idx2 := 0, 0

	for idx1 < le1 && idx2 < le2 {
		if str1[idx1] != str2[idx2] {
			if foundDiff {
				return false
			}

			foundDiff = true
		}

		idx1++
		idx2++
	}

	return true
}

// dpale
// dale
func oneEdit(str1, str2 string, le1, le2 int) bool {
	idx1, idx2 := 0, 0

	for idx1 < le1 && idx2 < le2 {
		if str1[idx1] != str2[idx2] {
			if idx1 != idx2 {
				return false
			}

			idx2++ // always assume string2 longer than string 1; skip this char -> all char affter is same
			continue
		}

		idx1++
		idx2++
	}

	return true
}

// alway str1 is shorter than str2
func merge2Case(str1, str2 string) bool {
	idx1, idx2 := 0, 0
	foundDiff := false
	for idx1 < len(str1) && idx2 < len(str2) {
		if str1[idx1] != str2[idx2] {
			if foundDiff {
				return false
			}

			foundDiff = true // skip the first diff
			// incase replace continue move short pointer
			if len(str1) == len(str2) {
				idx1++
			}
		} else {
			idx1++
		}

		idx2++
	}

	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
