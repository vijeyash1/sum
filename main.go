package add

func AddNumbers(a ...int) int {
	sum := 0
	for v := range a {
		sum += v

	}
	return sum
}
