package main

import (
	"fmt"
)

func Swap(sli []int, size int) {
	for i := 0; i < size-1; i++ {
		if sli[i] > sli[i+1] {
			sli[i], sli[i+1] = sli[i+1], sli[i]
		}
	}

}

func BubbleSort(unsorted []int) {
	size := len(unsorted)
	for i := 0; i < size; i++ {
		Swap(unsorted, size)
	}

}

func main() {
	var ele rune
	var size int
	var sli = make([]int, 10)
	size = len(sli)
	for i := 0; i <= size; i++ {
		if i >= len(sli) {
			size = size + 1
		}
		ele = 0
		fmt.Println("Введите число: ")
		fmt.Scan(&ele)
		if ele == 0 {
			fmt.Println("Ошибка")
			break
		}
		sli = append(sli, int(ele))
	}
	fmt.Println("До : ", sli)
	BubbleSort(sli)
	fmt.Println("После : ", sli)

}
