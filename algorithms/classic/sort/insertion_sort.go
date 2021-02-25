package sort

func InsertionSort(in []int) {
	for j := 1; j < len(in); j++ {
		key := in[j]
		i := j
		for i > 0 && in[i-1] > key {
			in[i] = in[i-1]
			i = i - 1
		}
		in[i] = key
	}
}
