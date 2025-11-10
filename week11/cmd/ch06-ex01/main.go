package main

import "fmt"

func main() {
	// var subjects []string
	// subjects = make([]string, 3)

	// subjects := make([]string, 3)
	// subjects[0] = "Go"
	// subjects[2] = "Python"

	subjects := []string{"Go", "", "Python"} // 슬라이스 리터럴

	for _, subject := range subjects {
		fmt.Println(subject)
	}
}
