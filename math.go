package main

func gcd(a, b int) int {
	for b > 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}
func gcdArray(input []int) int {
	result := input[0]
	for i := 1; i < len(input); i++ {
		result = gcd(result, input[i])
	}
	return result
}

func lcm(a, b int) int {
	return a * (b / gcd(a, b))
}

func lcmArray(input []int) int {
	result := input[0]
	for i := 1; i < len(input); i++ {
		result = lcm(result, input[i])
	}
	return result
}
