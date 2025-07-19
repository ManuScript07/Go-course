package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

var (
	DivisionByZeroError = errors.New("division by zero is not allowed")
	AmountIsIncorrect   = errors.New("amount is incorrect")
	BalanceIsIncorrect  = errors.New("balance is incorrect")
)

func PrintComplexNumber(z complex64) {
	realPart := real(z)
	imagPart := imag(z)

	roundedReal := math.Round(float64(realPart)*100) / 100
	roundedImag := math.Round(float64(imagPart)*100) / 100

	fmt.Printf("Действительная часть: %.2f. Мнимая часть: %.2f.\n", roundedReal, roundedImag)
}

func CheckLetters(text string) string {
	result := strings.Count(strings.ToLower(text), "е")
	if result == 0 {
		return "Текст готов к публикации!"
	}
	return fmt.Sprintf("Количество возможных ошибок: %d, перепроверьте текст", result)

}

var Balance float64 = 0

func topUpBalance(amount float64) error {
	if amount <= 0 {
		return AmountIsIncorrect
	}
	if Balance < 0 {
		return BalanceIsIncorrect
	}
	Balance += amount
	return nil
}

func chargeFromBalance(amount float64) error {
	if amount <= 0 {
		return AmountIsIncorrect
	}
	if Balance-amount < 0 {
		return BalanceIsIncorrect
	}
	Balance -= amount
	return nil
}

func TopUpAndGetBalance(amount float64) (float64, error) {
	err := topUpBalance(amount)
	if err != nil {
		return 0, err
	}
	return Balance, nil
}

func ChargeFromAndGetBalance(amount float64) (float64, error) {
	err := chargeFromBalance(amount)
	if err != nil {
		return 0, err
	}
	return Balance, nil
}

func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, DivisionByZeroError
	}
	return float64(a) / float64(b), nil
}

func task1() {
	var age int
	var name string
	var passportSeriesAndNumber string

	_, err := fmt.Scanln(&age, &name, &passportSeriesAndNumber)

	if err != nil {
		fmt.Println("Ошибка: некорректный ввод")
		return
	}
	if age < 14 || age > 150 {
		fmt.Println("Ошибка: невалидный возраст")
		return
	}
	if len(name) < 2 {
		fmt.Println("Ошибка: невалидное имя")
		return
	}
	if len(passportSeriesAndNumber) != 10 {
		fmt.Println("Ошибка: невалидная серия и номер паспорта")
		return
	}

	fmt.Println(fmt.Sprintf("Имя: %s. Возраст: %d. Серия и номер паспорта: %s", name, age, passportSeriesAndNumber))
}

func main() {
	task1()
}
