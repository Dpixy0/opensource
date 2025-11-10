// average calculates the average of several numbers.
package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

// go get 링크 주소로 먼저 적용시킴

func main() {
	numbers, err := datafile.GetFloats("meatWeight.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	weeks := float64(len(numbers))
	fmt.Printf("평균 고기 소비량 : %f", sum/weeks)
}
