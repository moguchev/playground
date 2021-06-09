package main

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	// _ "gopkg.in/goracle.v2"
)

func main() {
	now := time.Now()
	loc, _ := time.LoadLocation("Europe/Moscow")
	fmt.Println(now.UTC())
	fmt.Println(now.UTC().In(loc))
	t, _ := ptypes.TimestampProto(now.UTC())
	fmt.Println(t)
	t, _ = ptypes.TimestampProto(now.UTC().In(loc))
	fmt.Println(t)
}
