package main

import "fmt"

func PrettyArrayOutput(array [9]string) {
	for i, v := range array {
		if i < 7 {
			fmt.Printf("%d я уже сделал: %s\n", i+1, v)
			continue
		}
		fmt.Printf("%d не успел сделать: %s\n", i+1, v)
	}
}

func SumOfArray(array [6]int) int {
	sum := 0
	for _, i := range array {
		sum += i
	}
	return sum
}

func FindMinMaxInArray(array [10]int) (int, int) {
	max, min := array[0], array[0]
	for i := 0; i < 10; i++ {
		if max < array[i] {
			max = array[i]
		}
		if min > array[i] {
			min = array[i]
		}
	}
	return max, min
}

func ThirdElementI(array [6]int) int {
	return array[2]
}

func FiveSteps(array [5]int) [5]int {
	var t int
	for i := 0; i < 2; i++ {
		t = array[i]
		array[i] = array[4-i]
		array[4-i] = t
	}
	return array
}

func main() {
	var arr = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(SumOfArray(arr))
}
