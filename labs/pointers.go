package labs

import (
	"fmt"
)

// Срезы, указатели, прием множества значений в функцию, фор,
// инит, определние функции внутри мейн

var slcvar []int
var p *[]int

func init() {
	slcvar = []int{100000, 100, 10000, 10, 1000, 1000000}
	p = &slcvar
}

func incrementArr(slcvar *[]int) []int {
	// fmt.Println(*slcvar)
	for i := 0; i < len(*slcvar); i++ {
		(*slcvar)[i] *= i
	}
	fmt.Println(len(*slcvar), cap(*slcvar))
	fmt.Println(*slcvar)
	return (*slcvar)
}

func getNumsFindMin(numbers ...int) int {
	min := 0
	for _, num := range numbers {
		if num > min {
			min = num
		}
	}
	return min
}

func Pointers() {
	minfromnums := getNumsFindMin(incrementArr(p)...)
	fmt.Println(minfromnums)
	slcvar = append(slcvar, minfromnums)
	fmt.Println(len(slcvar), cap(slcvar))
	fmt.Println(*p)

	func() {
		incrementArr(p)
		// fmt.Println(slcvar)
	}()
}
