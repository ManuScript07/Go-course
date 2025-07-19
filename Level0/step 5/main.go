package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var startWork string
	var firstName, secondName, surname string
	var payment1, payment2, payment3 float64

	fmt.Scan(&startWork, &firstName, &secondName,
		&surname, &payment1, &payment2, &payment3)

	it, _ := time.Parse("02.01.2006", startWork)

	newDate := it.AddDate(0, 0, 15)

	totalPayment := payment1 + payment2 + payment3

	rubles := math.Floor(totalPayment)
	kopecks := math.Round((totalPayment - rubles) * 100)

	fmt.Printf(`Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.
Дата подписания договора: %s.
Просим вас подойти в офис в любое удобное для вас время в этот день.
Общая сумма выплат составит %.0f руб. %.0f коп.

С уважением,
Гл. бух. Иванов А.Е.`,
		firstName,
		secondName,
		surname,
		newDate.Format("02.01.2006"), // Форматирование даты для вывода
		rubles,
		kopecks,
	)
}
