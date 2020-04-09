package main

import (
	"fmt"
	"time"
)

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(arr[:3])
	fmt.Println(arr[3:3])
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
