package main

import (
	"fmt"
	"math"
)

func DeleteLongKeys(m map[string]int) map[string]int {
	for key, _ := range m {
		if len(key) < 5 {
			delete(m, key)
		}
	}
	return m
}

func CountingSort(contacts []string) map[string]int {

	m := make(map[string]int)

	for _, val := range contacts {
		m[val]++
	}
	return m
}

func SwapKeysAndValues(m map[string]string) map[string]string {
	newMap := make(map[string]string)

	for key, val := range m {
		newMap[val] = key
	}
	return newMap

}

func SumOfValuesInMap(m map[int]int) (s int) {
	for _, val := range m {
		s += val
	}
	return
}

func FindMaxKey(m map[int]int) int {
	mx := int(-math.Inf(1))
	for key, _ := range m {
		if key > mx {
			mx = key
		}
	}
	return mx
}

func main() {
	fmt.Println(FindMaxKey(map[int]int{1: 2, 3: 4}))
}
