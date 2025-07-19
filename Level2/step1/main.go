package main

import "fmt"

func BuyFries(size string) {
	switch size {
	case "S":
		printPrice(49, "Картошка фри")
	case "M":
		printPrice(89, "Картошка фри")
	case "L":
		printPrice(99, "Картошка фри")
	default:
		fmt.Println("Некорректный размер")
	}
}

func BuyCola(size string) {
	switch size {
	case "S":
		printPrice(99, "Кола")
	case "M":
		printPrice(119, "Кола")
	case "L":
		printPrice(139, "Кола")
	default:
		fmt.Println("Некорректный размер")
	}
}

func printPrice(price int, name string) {
	fmt.Printf("%s будет стоить %d рублей", name, price)
}

func PrintFlightRow(flightNumber, departureCity, arrivalCity string,
	duration float64, reg, gate int, ended bool) {
	if ended {
		fmt.Printf(`| %s | %s—%s | регистрация закончилась, проходите к гейту: %d | длительность полёта %.1f часа |
		`,
			flightNumber,
			departureCity,
			arrivalCity,
			gate,
			duration)
	} else {
		fmt.Printf(`| %s | %s—%s | %d регистрация продолжается |
		`,
			flightNumber,
			departureCity,
			arrivalCity,
			reg)
	}
}

func GoOrNot(ans string) {
	if ans == "Go" {
		fmt.Println("Go!")
		return
	}
	fmt.Println("Я знаю только Go!")
}

func main() {}
