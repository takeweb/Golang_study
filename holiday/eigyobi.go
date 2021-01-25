package main

import (
	"fmt"
	"time"

	"github.com/yut-kt/goholiday"
)

func main() {
	datetime := time.Date(2020, 9, 21, 0, 0, 0, 0, time.Local)

	// 休日か祝日の場合：true
	fmt.Println(goholiday.IsHoliday(datetime))

	// 祝日の場合：true
	fmt.Println(goholiday.IsNationalHoliday(datetime))

	// 営業日の場合：true
	fmt.Println(goholiday.IsBusinessDay(datetime))

	// 独自の休日を設定
	//var uniqueHolidayList []time.Time
	datetimeh1 := time.Date(2020, 12, 24, 0, 0, 0, 0, time.Local)
	datetimeh2 := time.Date(2020, 12, 25, 0, 0, 0, 0, time.Local)
	datetimeh3 := time.Date(2020, 12, 31, 0, 0, 0, 0, time.Local)
	// uniqueHolidayList = append(uniqueHolidayList, datetimeh1)
	// uniqueHolidayList = append(uniqueHolidayList, datetimeh2)
	// uniqueHolidayList = append(uniqueHolidayList, datetimeh3)
	uniqueHolidayList := []time.Time{datetimeh1, datetimeh2, datetimeh3}
	goholiday.SetUniqueHolidays(uniqueHolidayList)
	fmt.Println(goholiday.IsHoliday(datetimeh1))
}
