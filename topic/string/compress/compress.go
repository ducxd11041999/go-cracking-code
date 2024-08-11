package compress

import (
	"strconv"
)

func Compress(str string) string {

	leStr := len(str)
	compressed := ""
	same := 0
	for idx := 0; idx < leStr; idx++ {
		same++
		if idx+1 >= leStr || str[idx+1] != str[idx] {
			compressed += string(str[idx]) + strconv.Itoa(same)
			same = 0
		}
	}

	if len(compressed) < leStr {
		return compressed
	}

	return str
}

func CompressByte(str string) string {

	leStr := len(str)
	compressed := []byte{}
	same := 0
	for idx := 0; idx < leStr; idx++ {
		same++
		if idx+1 >= leStr || str[idx+1] != str[idx] {
			compressed = append(compressed, str[idx])
			compressed = append(compressed, byte(same+'0'))
			same = 0
		}
	}

	if len(compressed) < leStr {
		return string(compressed)
	}

	return str
}
