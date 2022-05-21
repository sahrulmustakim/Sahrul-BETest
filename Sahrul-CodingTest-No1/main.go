package main

import (
	"fmt"
	"math"
)

func convertBinaryToDecimal(number int) int {
	decimal := 0
	counter := 0.0
	remainder := 0

	for number != 0 {
		remainder = number % 10
		decimal += remainder * int(math.Pow(2.0, counter))
		number = number / 10
		counter++
	}
	return decimal
}

func convertDecimalToBinary(number int) int {
	binary := 0
	counter := 1
	remainder := 0

	for number != 0 {
		remainder = number % 2
		number = number / 2
		binary += remainder * counter
		counter *= 10
	}
	return binary
}

func main() {
	var binary int
	fmt.Print("Enter Binary Number : ")
	fmt.Scanln(&binary)

	fmt.Printf("Output : %d", convertBinaryToDecimal(binary))

	fmt.Println(" ")

	var decimal int
	fmt.Print("Enter Decimal Number : ")
	fmt.Scanln(&decimal)

	fmt.Printf("Output : %d", convertDecimalToBinary(decimal))
}
