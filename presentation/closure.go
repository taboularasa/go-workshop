package main

import "fmt"

func counterFactory(start, increment int) func() int {
	counter := start
	return func() int {
		counter = counter + increment
		return counter
	}
}

func main() {
	countBy10 := counterFactory(0, 10)
	fmt.Println(countBy10())
	fmt.Println(countBy10())
}
