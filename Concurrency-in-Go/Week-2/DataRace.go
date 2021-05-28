package main

import "fmt"

/*
Raca condition is when multiple threads are trying to access and manipulat the same variable.
In the code below, main, increment and decrement are all accessing and changing the value of x.
Due to the uncertainty of Goroutine scheduling mechanism, the results of the following program is unpredictable.
*/

func increment(a *int) {
	(*a)++
	fmt.Println(a)
}

func decrement(a *int) {
	(*a)--
	fmt.Println(a)
}

func main() {
	i := 0

	go increment(&i)
	go decrement(&i)

	fmt.Println(i)
}
