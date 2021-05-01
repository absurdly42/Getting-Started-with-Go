package main

import "fmt"

type Animals struct {
	food, locomotion, spoken string
}

func (w Animals) Eat() {
	fmt.Print(w.food)
}

func (w Animals) Move() {
	fmt.Print(w.locomotion)
}

func (w Animals) Speak() {
	fmt.Print(w.spoken)
}

func main() {
	m := map[string]Animals{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}
	for {
		fmt.Print(">")
		an := ""
		ac := ""
		fmt.Scan(&an, &ac)
		if ac == "eat" {
			m[an].Eat()
		} else if ac == "move" {
			m[an].Move()
		} else if ac == "speak" {
			m[an].Speak()
		}
	}
}
