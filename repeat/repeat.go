package repeat

func repeat(c, n int) int {
	for i := 0; i < n; i++ {
		c = c + i
	}
	return c
}
