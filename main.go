package main

import (
	"fmt"
	"main_go_in_4_hours/labs"
)

/*TODO Структуры + типы + конструкторы + интерфейсы*/

func main() {
	labs.NotDie()

	fmt.Println("Лаба TwoDimSlices")
	labs.TwoDimSlices()

	fmt.Println("Лаба Maps")
	labs.Maps()

	fmt.Println("Лаба Closure")
	labs.Closure()

	fmt.Println("Лаба AnonFunc")
	labs.AnonFunc()

	fmt.Println("Лаба FunctionsAndPrinting")
	labs.FunctionsAndPrinting()

	fmt.Println("Лаба MagicSlices")
	labs.MagicSlices()

	fmt.Println("Лаба Pointers")
	labs.Pointers()
}
