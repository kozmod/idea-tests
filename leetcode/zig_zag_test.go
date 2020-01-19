package leetcode

import (
	"testing"
)

func TestZigZag(t *testing.T) {
	s1 := "PAYPALISHIRING"
	rn1 := 3
	assert("PAHNAPLSIIGYIR", convert(s1, rn1), t)
}

func convert(s string, numRows int) string {
	sLength := len(s)
	if numRows < 2 || sLength <= numRows {
		return s
	}
	res := make([]byte, sLength)
	index := 0
	for i := 0; i < numRows; i++ {
		res[index] = s[i]
		index++
		tmp := i
		for tmp < sLength {
			tmp = tmp + 2*(numRows-1)
			if i != 0 && i != numRows-1 && tmp-2*i < sLength {
				res[index] = s[tmp-2*i]
				index++
			}
			if tmp < sLength {
				res[index] = s[tmp]
				index++
			}
		}
	}
	return string(res)
}

func assert(ex string, r string, t *testing.T) {
	if ex != r {
		t.Errorf("convert = %s; want = %s", r, ex)
	}
}
