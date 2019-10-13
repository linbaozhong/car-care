package utils

import (
	"bytes"
)

func ToHex(n int64) string {
	h := [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	s := new(bytes.Buffer)
	for ; n > 0; n /= 16 {
		s.WriteString(h[n%16])
	}
	return string(reverse(s.Bytes()))
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
