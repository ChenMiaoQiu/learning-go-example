package popcount4

func lowbit(x uint64) uint64 {
	return x & -x
}

func PopCount(x uint64) int {
	res := 0
	for x != 0 {
		res++
		x -= lowbit(x)
	}
	return res
}
