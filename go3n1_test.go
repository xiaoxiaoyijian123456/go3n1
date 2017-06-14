package go3n1

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCnt3n1(t *testing.T) {
	tests := [][2]uint64{
		{1, 1}, {2, 2}, {3, 8}, {4, 3}, {5, 6}, {6, 9}, {7, 17}, {8, 4}, {9, 20}, {10, 7},
	}
	for _, tt := range tests {
		if ret := Cnt3n1(tt[0]); ret != tt[1] {
			t.Errorf("n = %d, ret = %d, not match: %d", tt[0], ret, tt[1])
		}
	}
}

func TestList3n1(t *testing.T) {
	tests := []struct {
		N    uint64
		List []uint64
	}{
		{N: 1, List: []uint64{1}},
		{N: 2, List: []uint64{2, 1}},
		{N: 3, List: []uint64{3, 10, 5, 16, 8, 4, 2, 1}},
		{N: 4, List: []uint64{4, 2, 1}},
		{N: 5, List: []uint64{5, 16, 8, 4, 2, 1}},
		{N: 6, List: []uint64{6, 3, 10, 5, 16, 8, 4, 2, 1}},
		{N: 7, List: []uint64{7, 22, 11, 34, 17, 52, 26, 13, 40, 20, 10, 5, 16, 8, 4, 2, 1}},
		{N: 8, List: []uint64{8, 4, 2, 1}},
		{N: 9, List: []uint64{9, 28, 14, 7, 22, 11, 34, 17, 52, 26, 13, 40, 20, 10, 5, 16, 8, 4, 2, 1}},
		{N: 10, List: []uint64{10, 5, 16, 8, 4, 2, 1}},
	}
	for _, tt := range tests {
		retChan := List3n1(tt.N)
		ret := []uint64{}
		for v := range retChan {
			ret = append(ret, v)
		}
		if !uint64_slice_equals(ret, tt.List) {
			t.Errorf("n = %d, ret = %v, not match: %v", tt.N, uint64_slice_str(ret), uint64_slice_str(tt.List))
		}
	}
}

func TestMaxlen3n1(t *testing.T) {
	tests := []struct {
		M      uint64
		N      uint64
		MaxNum uint64
		MaxLen uint64
	}{
		{M: 1, N: 10, MaxNum: 9, MaxLen: 20},
		{M: 20, N: 100, MaxNum: 97, MaxLen: 119},
		{M: 1000, N: 5000, MaxNum: 3711, MaxLen: 238},
		{M: 10000, N: 50000, MaxNum: 35655, MaxLen: 324},
	}
	for _, tt := range tests {
		if ret, maxlen := Maxlen3n1(tt.M, tt.N); ret != tt.MaxNum || maxlen != tt.MaxLen {
			t.Errorf("m = %d, n = %d, (ret, maxlen) = (%d, %d), not match: (%d, %d)", tt.M, tt.N, ret, maxlen, tt.MaxNum, tt.MaxLen)
		}
	}
}

func uint64_slice_equals(a, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func uint64_slice_str(a []uint64) string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i, v := range a {
		if i == 0 {
			buffer.WriteString(fmt.Sprintf("%d", v))
		} else {
			buffer.WriteString(fmt.Sprintf(", %d", v))
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}
