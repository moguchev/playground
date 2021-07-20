package main

import (
	"fmt"
	"time"
	// "github.com/golang/protobuf/ptypes"
	// _ "gopkg.in/goracle.v2"
)

func main() {
	a, b := durationToHoursAndMinutes(time.Duration(43200+1800) * time.Second)

	fmt.Println(a, ":", b)
}

func durationToHoursAndMinutes(d time.Duration) (hours, minutes int) {
	hours = int(d / time.Hour)
	minutes = int((d % time.Hour) / time.Minute)
	return
}
