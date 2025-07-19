package main

import "fmt"

func task1() {
	fmt.Println("Я люблю Go!")
}

func task2() {
	var favthing string
	fmt.Scanln(&favthing)

	fmt.Println(
		fmt.Sprintf("А ещё я люблю %s!", favthing),
	)
}

func task3() {
	var name string
	var apart, password int
	var duration float64
	fmt.Scan(&name, &apart, &password, &duration)

	fmt.Println(
		fmt.Sprintf(
			"Привет, %s! Приглашаю тебя на соревнование по программированию, которое пройдёт, как всегда, в квартире %d. Оно будет длиться примерно %.1f часа. Не забудь секретный пароль для входа: %d.",
			name,
			apart,
			duration,
			password,
		),
	)

}

func main() {
	// task1()
	// task2()
	task3()
}
