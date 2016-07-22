package control_structures

import "fmt"

func fizzBuzz(i int32) string {

	if i%3!=0 && i%5!=0{
		return fmt.Sprintf("%d",i)
	}else if i%3==0 && i%5==0{
		return "FizzBuzz"
	}else if i%3==0 {
		return "Fizz"
	}else if i%5==0 {
		return "Buzz"
	}
	return ""
}
