package main

import (
	"fmt"
)

func main() {
	fizzBuzz(15)
}


func fizzBuzz(n int32) {
	var str string
	for i := int32(1); i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			str = "FizzBuzz"
		} else if i%3 == 0 {
			str = "Fizz"
		} else if i%5 == 0 {
			str = "Buzz"
		} else {
			str = fmt.Sprintf("%v", i)
		}
		fmt.Println(str)
	}
}