package main

import "fmt"

func main() {
	subjects := [4]string{"Go", "Javascript", "Python", "Linux"} // 슬라이스 리터럴
	subjectsSlice := subjects[1:3]                               // 슬라이싱
	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("================")
	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
}
