package main

import "fmt"

func main() {
	var x float32
	fmt.Print("Введите дробное число: ")
	fmt.Scan(&x)
	fmt.Println(int(x))
}
