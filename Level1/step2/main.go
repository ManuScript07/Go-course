package main

import (
	"fmt"
)

func task1() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	if a == b && b == c {
		fmt.Println("Максимальное равенство")
	} else {
		fmt.Println("Не равны")
	}
}

func task2() {
	var pass1, pass2 string
	fmt.Scan(&pass1, &pass2)
	len1 := len(pass1)
	len2 := len(pass2)
	if len1 > 7 && len2 > 7 {
		fmt.Println("Оба пароля надёжные")
	} else if len1 > 7 && len2 < 8 {
		fmt.Println("Только первый пароль надёжный")
	} else if len1 < 8 && len2 > 7 {
		fmt.Println("Только второй пароль надёжный")
	} else {
		fmt.Println("Оба пароля ненадёжные")
	}
}

func task3() {
	var a, b string
	fmt.Scan(&a, &b)
	if a == b {
		fmt.Println("Ничья")
	} else if a == "камень" {
		if b == "ножницы" {
			fmt.Println("Первый игрок победил")
		} else {
			fmt.Println("Второй игрок победил")
		}
	} else if a == "ножницы" {
		if b == "бумага" {
			fmt.Println("Первый игрок победил")
		} else {
			fmt.Println("Второй игрок победил")
		}
	} else if a == "бумага" {
		if b == "камень" {
			fmt.Println("Первый игрок победил")
		} else {
			fmt.Println("Второй игрок победил")
		}
	}
}

func task4() {
	var sign string
	var degree int
	fmt.Scanln(&sign, &degree)
	if sign == "+" && degree > 20 {
		fmt.Println("Стоит надеть майку и шорты")
	} else if sign == "+" && degree >= 10 {
		fmt.Println("Стоит надеть штаны и кофту")
	} else if sign == "-" && degree > 5 {
		fmt.Println("Стоит надеть зимнюю куртку")
	} else {
		fmt.Println("Стоит надеть куртку")
	}
}

func main() {
	// task1()
	// task2()
	// task3()
	task4()
}
