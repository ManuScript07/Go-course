package main

import "fmt"

func main() {
	a, b, c, d, e := "-", "-", "-", "-", "-"
	var input string
	var freePlaces int = 0

	for {
		fmt.Scan(&input)
		switch input {
		case "очередь", "конец":
			for i := 0; i < 5; i++ {
				switch i {
				case 0:
					fmt.Println("1. ", a)
				case 1:
					fmt.Println("2. ", b)
				case 2:
					fmt.Println("3. ", c)
				case 3:
					fmt.Println("4. ", d)
				case 4:
					fmt.Println("5. ", e)
				}
			}
			if input == "конец" {
				return
			}
		case "количество":
			fmt.Println("Осталось свободных мест: ", freePlaces)
			fmt.Println("Всего человек в очереди: ", 5-freePlaces)

		}
	}
}
