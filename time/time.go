package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	day := 24 * time.Hour
	transferDate := now.Add(-3 * time.Hour).Unix()
	fmt.Println(now, transferDate)
	if time.Unix(transferDate, 0).Truncate(day).Equal(now.Truncate(day)) {
		fmt.Println("EQUAL DATES")
	}
}
