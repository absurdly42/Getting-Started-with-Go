package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, addr string
	var person map[string]string
	person = make(map[string]string)
	fmt.Print("Введите имя и адрес: ")
	fmt.Scan(&name, &addr)
	person[name] = addr
	fmt.Println(person)
	barr, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Ошибка")
	}
	fmt.Println(barr)

}
