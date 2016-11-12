package main

import (
	"fmt"
)

const (
	Satisfied = "satisfied"
	Hungry    = "hungry"

	Inconsistent = "Inconsistent"

	MinWalnuts = 2
	MaxWalnuts = 10
)

func main() {
	var lines int
	fmt.Scanf("%d\n", &lines)

	satCount := MaxWalnuts
	hungCount := MinWalnuts
	var walnuts int
	var status string

	for i := 0; i < lines; i++ {
		fmt.Scanf("%d %s\n", &walnuts, &status)

		switch status {
		case Satisfied:
			if walnuts < satCount {
				satCount = walnuts
			}
		case Hungry:
			if walnuts > hungCount {
				hungCount = walnuts
			}
		}
	}

	if satCount <= hungCount {
		fmt.Println(Inconsistent)
	} else {
		fmt.Println(satCount)
	}
}
