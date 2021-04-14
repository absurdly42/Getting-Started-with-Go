package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Scan(&x)
	n := len(x)
	r := x[n/2 : n/2+1]
	x = strings.ToLower(x)
	if strings.HasPrefix(x, "i") && strings.HasSuffix(x, "n") && r == "a" {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
