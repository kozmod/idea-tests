package search

type testCase struct {
	in  []int
	exp int
}

func reqMax(in []int) int {
	if len(in) < 2 {
		return in[0]
	}
	l := reqMax(in[:len(in)/2])
	r := reqMax(in[len(in)/2:])
	if l > r {
		return l
	}
	return r
}

func linearMax(in []int) int {
	var max *int
	for _, v := range in {
		if max == nil || *max < v {
			max = &v
		}
	}
	if max == nil {
		return 0
	}
	return *max
}
