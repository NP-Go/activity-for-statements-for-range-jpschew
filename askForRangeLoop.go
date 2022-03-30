package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Insert code for only even here
	var number1, number2 string
	var odd, even []int
	fmt.Println("Please enter first number: ")
	fmt.Scanln(&number1)
	fmt.Println("Please enter second number: ")
	fmt.Scanln(&number2)
	num1, _ := strconv.ParseInt(number1, 10, 0)
	num2, _ := strconv.ParseInt(number2, 10, 0)

	if num1 > num2 {
		num1, num2 = num2, num1
	}

	for i := num1; i <= num2; i++ {
		if num1 == num2 {
			break
		}
		if i%2 == 0 {
			even = append(even, int(i))
		} else {
			odd = append(odd, int(i))
		}
	}

	fmt.Printf("The even numbers between %d and %d are %v\n", num1, num2, even)
	fmt.Printf("The odd numbers between %d and %d are %v\n", num1, num2, odd)
}
