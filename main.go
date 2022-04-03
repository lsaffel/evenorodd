package main

import "fmt"

func main()  {

	// create a slice of integers from 0 through 10
	var intSlice []int

	for i := 0; i<=10; i++ {
		intSlice = append(intSlice, i)
	}

	fmt.Println(intSlice)

	// for each number in the int slice, if it's even, 
	// print that it is even. Otherwise, print that it's odd
	for _,num := range intSlice {
		if num % 2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}	
}