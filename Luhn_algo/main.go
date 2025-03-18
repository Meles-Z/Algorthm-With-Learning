package main

import "fmt"

func LuhanAlgorithm(num int) bool {
	sum := 0
	double := false

	for num != 0 {
		digit := num % 10
		num /= 10

		if double {
			digit *= 2
			if digit > 9 {
				digit = digit%10 + digit/10 // Sum the two digits directly
			}
		}

		sum += digit
		double = !double // Toggle the doubling for every second digit
	}

	fmt.Println(sum)
	return sum%10 == 0
}

func main() {
	fmt.Println(LuhanAlgorithm(79927398713))       // Should return true
	fmt.Println(LuhanAlgorithm(4532015112830366))  // Should return true
}

