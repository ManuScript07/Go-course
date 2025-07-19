package main

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	dayError = errors.New("исправь свой ответ, а лучше ложись поспать")
)

func TimeNow() time.Time

func currentDayOfTheWeek() string {
	timeNow := TimeNow()
	nDay := int(timeNow.Weekday()) - 1

	var weekDay string

	switch nDay {
	case 0:
		weekDay = "Понедельник"
	case 1:
		weekDay = "Вторник"
	case 2:
		weekDay = "Среда"
	case 3:
		weekDay = "Четверг"
	case 4:
		weekDay = "Пятница"
	case 5:
		weekDay = "Суббота"
	case 6:
		weekDay = "Воскресенье"
	}
	return weekDay
}

func dayOrNight() string {
	timeNow := TimeNow()
	hour := timeNow.Hour()
	if hour >= 10 && hour <= 22 {
		return "День"
	}
	return "Ночь"
}

func nextFriday() int {
	timeNow := int(TimeNow().Weekday()) - 1
	if timeNow == 4 {
		return 0
	} else if timeNow < 4 {
		return 4 - timeNow
	} else {
		return 7 - timeNow + 5
	}

}

func CheckCurrentDayOfTheWeek(answer string) bool {
	answer = strings.ToLower(answer)
	return answer == strings.ToLower(currentDayOfTheWeek())
}

func CheckNowDayOrNight(answer string) (bool, error) {
	answer = strings.ToLower(answer)
	if utf8.RuneCountInString(answer) != 4 {
		return false, dayError
	}
	return answer == strings.ToLower(dayOrNight()), nil
}
