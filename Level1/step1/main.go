package main

import (
	"fmt"
	"math"
)

func task1() {
	var input string
	fmt.Scanln(&input)

	if input == "Go" {
		fmt.Println("Go!")
		return
	}
	fmt.Println("Я знаю только Go!")
}

func task2() {
	var a, b int
	fmt.Scanln(&a, &b)

	if a > b {
		fmt.Println("Первое число больше второго")
	} else if a < b {
		fmt.Println("Второе число больше первого")
	} else {
		fmt.Println("Числа равны")
	}

}

func task3() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)
	if a+b/100 >= c {
		fmt.Println("Сегодня будет вкусный кофе!")
	} else {
		fmt.Println("Стоит подкопить")
	}
}

func task4() {
	var n float64
	fmt.Scan(&n)
	if n >= 0 {
		fmt.Println(fmt.Sprintf("%.3f", math.Sqrt(n)))
	} else {
		fmt.Println(-1)
	}
}

func task5() {
	var n int
	fmt.Scanln(&n)
	if n == 0 {
		fmt.Println("Число 0")
	} else if -10 < n && n < 10 {
		fmt.Println("Число однозначное")
	} else if n%2 == 0 {
		fmt.Println("Число чётное")
	} else if n > 0 {
		fmt.Println("Число положительное")
	} else {
		fmt.Println("Число красивое")
	}
}

func main() {
	// task1()
	// task2()
	// task3()
	// task4()
	task5()
}
