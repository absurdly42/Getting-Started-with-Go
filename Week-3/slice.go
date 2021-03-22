package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	List := make([]int, 0, 3)

	var scan string
	for {
		fmt.Print("Введите число:")
		_, E := fmt.Scan(&scan)
		exit := strings.ToLower(scan)
		if E != nil {
			fmt.Println(E)
			os.Exit(0)
		}
		if exit == "x" {
			fmt.Println("Завершение:")
			os.Exit(0)
		}

		parseIntScanValue, E := strconv.Atoi(scan)
		List = append(List, parseIntScanValue)
		sort.Ints(List)
		fmt.Println(List)
	}
}
