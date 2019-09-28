package main

import "fmt"

//Fizz Buzz - Write a program that prints the numbers from 1 to 100.
//But for multiples of three print “Fizz” instead of the number and
//for the multiples of five print “Buzz”.
//For numbers which are multiples of both three and five print “FizzBuzz”.

func main() {

	for i := 0; i < 100; i++ {

		// Start with most complicated first
		// Multiple of 3 & 5
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		}

	}

}
