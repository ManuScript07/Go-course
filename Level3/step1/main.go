package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func NumbersToLetters(input string) string {
	var output string
	output = strings.Replace(input, "0", "ноль", -1)
	output = strings.Replace(output, "1", "один", -1)
	output = strings.Replace(output, "2", "два", -1)
	output = strings.Replace(output, "3", "три", -1)
	output = strings.Replace(output, "4", "четыре", -1)
	output = strings.Replace(output, "5", "пять", -1)
	output = strings.Replace(output, "6", "шесть", -1)
	output = strings.Replace(output, "7", "семь", -1)
	output = strings.Replace(output, "8", "восемь", -1)
	output = strings.Replace(output, "9", "девять", -1)
	output = strings.Replace(output, "+", "плюс", -1)
	output = strings.Replace(output, "-", "минус", -1)
	output = strings.Replace(output, "*", "умножить на", -1)
	output = strings.Replace(output, "/", "разделить на", -1)
	return output
}

func CheckOnlyASCII(s string) bool {
	for _, letter := range s {
		if letter > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func CountLengthAndBytes(first, second string) string {
	single := first + second
	return fmt.Sprintf(`Объединённая строка: %s. Количество байт: %d. Количество символов: %d.`,
		single,
		len(single),
		utf8.RuneCountInString(single))
}
