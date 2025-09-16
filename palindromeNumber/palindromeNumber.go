package main

import "fmt"

func main() {
	fmt.Println(palindromeNumber(323))
}

// currently not working
func palindromeNumber(x int) bool {
	if x < 0 {
		return false
	}
	num := x
	reversed := 0
	for x != 0 {
		reversed = 10*reversed + num%10
		x /= 10
	}
	return (reversed == x)
}
