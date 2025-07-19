package main

import (
	"fmt"
	"strings"
)

func task1() {
	var ans string

	for i := 0; i < 5; i++ {
		fmt.Scan(&ans)
		if ans == "Go" {
			fmt.Println("Go!")
		} else {
			fmt.Println("Я знаю только Go!")
		}
	}
}

func task2() {
	var N int
	var mark float64

	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		fmt.Scanln(&mark)
		switch {
		case mark >= 90 && mark <= 100:
			fmt.Println(5)
		case mark >= 75 && mark <= 89:
			fmt.Println(4)
		case mark >= 50 && mark <= 74:
			fmt.Println(3)
		case mark >= 0 && mark <= 49:
			fmt.Println(2)
		default:
			fmt.Println("Неверный балл")
		}
	}
}

func task3() {
	var N, discount int
	var price float64
	var sum float64 = 0

	fmt.Scan(&N, &discount)

	k := 1.0 - float64(discount)/100.0

	for i := 0; i < N; i++ {
		fmt.Scan(&price)
		sum += price * k
	}

	fmt.Println(sum)

}

func task4() {
	var ans string
	var end bool = false
	for {
		fmt.Scan(&ans)
		switch ans {
		case "да", "нет", "чёрный", "белый":
			fmt.Println("Поражение")
			end = true
		default:
			fmt.Println("Игра продолжается")
		}
		if end {
			break
		}
	}
}

func task5() {
	var N, k int
	var word string
	var w string
	fmt.Scan(&N, &word)

	for i := 0; i < N; i++ {
		fmt.Scan(&w)
		if strings.ToLower(w) == strings.ToLower(word) {
			k += 1
		}
	}
	fmt.Println(k)
}

func task6() {
	var word string
	fmt.Scanln(&word)

	for _, letter := range word {
		if letter == 'а' || letter == 'о' {
			continue
		}
		fmt.Print(string(letter))
	}
}

func main() {
	// task1()
	// task2()
	// task3()
	// task4()
	// task5()
	task6()
}
