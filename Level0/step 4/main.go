package main

import (
	"fmt"
	"math"
	"time"
)

func task1() {
	var height1, height2, height3 float64
	fmt.Scan(&height1, &height2, &height3)
	fmt.Println(math.Min(math.Min(height1, height2), height3))
}

func task2() {
	var x, y, z, m, n float64
	fmt.Scan(&x, &y, &z, &m, &n)
	result := (5 * x * math.Sin(2*y)) / (z + math.Pow(n, math.Log(m)))
	fmt.Println(result)
}

func task3() {
	var a, b float64
	fmt.Scan(&a, &b)
	result := math.Round(math.Max(a, b))
	fmt.Println(result)
}

func task4() {
	var timeNow string
	fmt.Scanln(&timeNow)
	var format string = "2006-01-02/15:04:05"
	t, _ := time.Parse(format, timeNow)
	fmt.Printf("Текущее время %d часов, %d минут. Ты точно не забыл про важные дела на сегодня?",
		t.Hour(),
		t.Minute(),
	)
}

func task5() {
	var time1, time2 string
	fmt.Scan(&time1, &time2)
	var format string = "2006-01-02"
	t1, _ := time.Parse(format, time1)
	t2, _ := time.Parse(format, time2)
	diff := t1.Year() - t2.Year()
	fmt.Println(diff, "yers ago")

}

func main() {
	// task1()
	// task2()
	// task3()
	// task4()
	task5()
}
