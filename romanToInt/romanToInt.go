package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
}

func romanToInt(s string) int {
	romanNums := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	number := 0
	for _, v := range s {
		number += romanNums[string(v)]
	}

	return number
}
