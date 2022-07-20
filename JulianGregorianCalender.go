package main

import (
	"fmt"
	"strconv"
)

func dayOfProgrammer(year int) string {
	totalDays := 215 //exclude feb
	programmerDays := 256
	if year < 1919 {
		if year%4 == 0 {
			totalDays += 29
		} else {
			totalDays += 28
		}
	} else {
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			totalDays += 29
		} else {
			totalDays += 28
		}
	}
	data := fmt.Sprintf("%s.09.%s", strconv.Itoa(programmerDays-totalDays), strconv.Itoa(year))
	return data
}

func main() {
	date := dayOfProgrammer(2017)
	fmt.Print(date)
}
