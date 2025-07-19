package main

import (
	"math"
	"unicode"
)

func ratePassword(password string) string {
	var rate int = 0
	if hasMinimumLength(password, 8) {
		rate++
	}
	if hasUpper(password) {
		rate++
	}
	if hasLowerCase(password) {
		rate++
	}

	switch rate {
	case 3:
		return "Сложный пароль"
	case 2:
		return "Средний пароль"
	case 1:
		return "Слабый пароль"
	default:
		return "Пароль недостаточно безопасен, придумайте новый"
	}
}

func hasMinimumLength(password string, lenght int) bool {
	return len(password) >= lenght
}

func hasUpper(password string) bool {
	for _, letter := range password {
		if unicode.IsUpper(letter) {
			return true
		}
	}
	return false
}

func hasLowerCase(password string) bool {
	for _, letter := range password {
		if unicode.IsLower(letter) {
			return true
		}
	}
	return false
}

func SquareRoots(a, b, c float64) (float64, float64) {
	D := findDiscriminant(a, b, c)
	if D < 0 {
		return 0, 0
	} else if D == 0 {
		result := -b / (2 * a)
		return result, result
	}
	x1 := (-b + math.Sqrt(D)) / (2 * a)
	x2 := (-b - math.Sqrt(D)) / (2 * a)
	return x2, x1
}

func findDiscriminant(a, b, c float64) float64 {
	return b*b - 4*a*c
}

func main() {}
