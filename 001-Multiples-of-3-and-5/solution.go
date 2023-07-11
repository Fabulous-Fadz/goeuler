package multiplesof3and5

func RunBrute(upToExclusive int) int {
	// If it's a small set so we can brute force...
	var sum int
	for i := 1; i < upToExclusive; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func totalDivisibleBy(input, target int) int {
	res := (target - 1) / input
	return input * (res * (res + 1)) / 2
}

func RunSmart(upTo int) int {
	return totalDivisibleBy(3, upTo) + totalDivisibleBy(5, upTo) - totalDivisibleBy(15, upTo)
}
