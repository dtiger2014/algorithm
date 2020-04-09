package main

import (
	"fmt"
	"time"
)

func main() {
	getMonday()
}

func getMonday() {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	fmt.Println(weekStartDate.Unix())
}
