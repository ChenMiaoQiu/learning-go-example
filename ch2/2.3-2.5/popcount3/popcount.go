package popcount3

func PopCount(x uint64) int {
	res := 0
	for x != 0 {
		if x&1 == 1 {
			res++
		}
		x >>= 1
	}
	return res
}
