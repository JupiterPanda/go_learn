package labs

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func NotDie() {
	fmt.Println("")
}

// Двумерные массивы, обработка panic, defer
func TwoDimSlices() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Произошла ошибка:", err)
		}
	}()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var n, m, summdiagmain, summdiagsec, max int
	fmt.Print("Enter n, m: ")
	fmt.Scan(&n, &m)
	matr := make([][]int, n)
	summcol := make([]int, m)
	summstr := make([]int, n)
	quad := (n == m)

	for i := 0; i < n; i++ { // фор по столбцу
		max = 0
		matr[i] = make([]int, m) // создаем строку длинной m
		for j := 0; j < m; j++ { // фор по строке
			matr[i][j] = r.Intn(10) // присваиваем элементу рандомное значение

			summstr[i] += matr[i][j] // подсчёт суммы для каждой строки
			summcol[j] += matr[i][j] // подсчёт суммы для каждого столбца

			if max < matr[i][j] {
				max = matr[i][j]
			}
			if quad {
				if i == j {
					summdiagmain += matr[i][j]
				}
				if i == (m - j) {
					summdiagsec += matr[i][j]
				}
			}
		}
		// вывод строки & максимума в строке
		fmt.Println(matr[i], "Максимум в строке", max)
	}
	fmt.Println("Сумма по столбцам: ", summcol, "Сумма по строкам: ", summstr)
	if quad {
		fmt.Println("Сумма по главной диагонали: ", summdiagmain,
			"Сумма по побочной диагонали: ", summdiagsec)
	}
}

// Мапы
func Maps() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Произошла ошибка:", err)
		}
	}()

	selection := map[string]int{"Jeff": 0, "Bob": 0, "Paul": 0}
	max_votes := 0
	var winner string

	for i := 0; i < 7; i++ {
		var vote string
		fmt.Scan(&vote)
		_, ok := selection[vote]
		if ok {
			selection[vote]++
			fmt.Println(selection)
		} else {
			fmt.Println("Нет такого кандидата или введены некорректные данные - попробуйте снова")
			i--
		}
	}
	for candidate, num_of_votes := range selection {
		if num_of_votes > max_votes {
			max_votes = num_of_votes
			winner = candidate
		}
	}
	deleteLosers(&selection)
	fmt.Println("FUCK losers: ", selection)

	fmt.Println("Winner: ", winner, max_votes)

}

func deleteLosers(selection *map[string]int) {

	for candidate, num_of_votes := range *selection {
		if num_of_votes < 2 {
			delete(*selection, candidate)
		}
	}
}

// Замыкание
func Closure() {
	inc := increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}

func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// анонимные функции и входное значение в виде функции

func testSumm(summfunc func(int, int) int) {
	fmt.Println(summfunc(3, 4))
}

func AnonFunc() {
	summ := func(a int, b int) int { return a + b }
	testSumm(summ)
}

// Функции и выводы
// Is float a higher than complex of b?

func FunctionsAndPrinting() {

	a := 23.2
	b := 23 + 6i

	res, err := higherlowercase(a, b)
	if err != nil {
		log.Fatal(err)
	}
	if res {
		fmt.Printf("A is higher than B, foo returned %t", res)
	} else {
		fmt.Printf("B is higher than A, foo returned %t", res)
	}
}

// на ифах
// func higherlowerif(a float64, b complex128) (bool, error) {
// 	if a > real(b) {
// 		return true, nil
// 	} else if a < real(b) {
// 		return false, nil
// 	} else {
// 		return false, errors.New("they're equal")
// 	}
// }

// на свитчах
func higherlowercase(a float64, b complex128) (bool, error) {
	switch {
	case a > real(b):
		return true, nil
	case a < real(b):
		return false, nil
	case a == real(b):
		return false, errors.New("they're equal")
	default:
		return false, errors.New("error")
	}
}

// Колдовство со слайсами
func MagicSlices() {
	sample := make([]int, 0, 5)

	for i := 0; i < 5; i++ {
		sample = append(sample, i)
	}

	fmt.Println(len(sample), cap(sample))
	sample = sample[:0]

	// fmt.Println(len(sample), cap(sample))

	fmt.Println(sample[:])

	fmt.Println(sample[4])
	fmt.Println(sample[:3])
	fmt.Println(sample[4])
	fmt.Println(sample[:4])
	fmt.Println(sample[4])

}
