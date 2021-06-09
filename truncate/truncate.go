package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {

	value := 0.07
	value *= 100

	v, _ := decimal.NewFromFloat(value).Round(1).Float64()
	value = v
	fmt.Printf("%f", v)
}
