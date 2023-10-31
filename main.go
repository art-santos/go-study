//This code prints out the first 100 Fibonacci numbers, along with their English name. It uses the Fibonacci function from the math/fibonacci package, and the NumberToWords function from the math/naming package. If there is an error converting a number to words, it prints out an error message instead of the number.

package main

import (
	"fmt"
	"study/src/math/fibonacci"
	"study/src/math/naming"
)

func main() {
    n := 100
    seq := fibonacci.Fibonacci(n)


for _, number := range seq {
		englishName, err := naming.NumberToWords(number)
		if err != nil {
			fmt.Printf("Error converting %d: %v\n", number, err)
		} else {
			fmt.Printf("%d in English is: %s\n", number, englishName)
		}
	}
}