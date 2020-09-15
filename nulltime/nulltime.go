package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
)

const (
	OracleNullDate = "01.01.0001"
	DateLayout     = "02.01.2006"
)

type SalaryInfo struct {
	DateFrom time.Time
}

func main() {
	v := SalaryInfo{}
	var realChangeDate sql.NullTime
	realChangeDate.Scan(nil)

	if realChangeDate.Time.Format(DateLayout) != OracleNullDate {
		fmt.Println("realChangeDate.Time", realChangeDate.Time)
		v.DateFrom = realChangeDate.Time
	} else {
		fmt.Println("NULL", realChangeDate.Time)
	}

	fmt.Println("v.DateFrom: ", v.DateFrom)
	dateFrom, err := ptypes.TimestampProto(v.DateFrom)
	if err != nil {
		fmt.Println("can't convert date from to timestamp")
	}
	fmt.Println("dateFrom: ", dateFrom)
}
