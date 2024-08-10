package unique

func IsUniqueElementUsingArray(str string) bool {
	le := len(str)
	unique := make([]bool, 128)
	for i := 0; i < le; i++ {
		if unique[byte(str[i])] {
			return false
		}

		unique[byte(str[i])] = true
	}

	return true
}

func IsUniqueElementUsingMap(str string) bool {
	le := len(str)
	unique := map[byte]bool{}
	for i := 0; i < le; i++ {
		if _, ok := unique[byte(str[i])]; ok {
			return false
		}

		unique[byte(str[i])] = true
	}

	return true
}

func IsUniqueElementUsingBitsMap(str string) bool {
	le := len(str)
	unique := make([]int64, 4) // 4 * 64 bits = 256 bits

	for i := 0; i < le; i++ {
		val := int(str[i])
		bitIndex := val / 64
		bitPosition := val % 64

		if unique[bitIndex]&(1<<bitPosition) > 0 {
			return false
		}

		unique[bitIndex] |= 1 << bitPosition
	}

	return true
}
