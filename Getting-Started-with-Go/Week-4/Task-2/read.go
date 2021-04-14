package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var name []Name
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, _, err := rd.ReadLine()

		if err != nil || io.EOF == err {
			break
		}
		tname := strings.Split(string(line), " ")
		tpname := Name{
			LongName(tname[0]),
			LongName(tname[1]),
		}
		name = append(name, tpname)
	}
	for i := 0; i < len(name); i++ {
		fmt.Println(i + 1)
		fmt.Println("First Name:" + name[i].fname)
		fmt.Println("Last Name:" + name[i].lname)
	}
}

func LongName(buffer string) string {

	if len(buffer) > 20 {
		return string(buffer[0:20])
	} else {
		return buffer
	}
}
