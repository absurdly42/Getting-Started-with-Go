package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sli := make([]int, 0, 3)

	var scan string
	fmt.Print("Введите число: ")
	for {
		fmt.Scan(&scan)
		exit := strings.ToLower(scan)
		if exit == "x" {
			fmt.Println("Завершение программы")
			os.Exit(0)
		} else {
			str, err := strconv.Atoi(scan)
			if err != nil {
				fmt.Println("Повторите попытку")
			}
			sli = append(sli, str)
			sort.Ints(sli)
			fmt.Println(sli)
		}

	}
}
