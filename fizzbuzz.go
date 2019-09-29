package main

import "fmt"

/*
The original classic overused initial job interview question

Fizz Buzz - Write a program that prints the numbers from 1 to 100.
for multiples of three print “Fizz” instead of the numver and
for the multiples of five print “Buzz”
for numbers which are multiples of both three and five print “FizzBuzz”.
*/

func main() {

	for i := 0; i < 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

	}

}
