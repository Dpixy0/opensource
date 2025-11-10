package main

import "fmt"

func main() {
	subjects := [4]string{"Go", "Javascript", "Python", "Linux"} // 슬라이스 리터럴
	subjectsSlice := subjects[:3]
	subjects[0] = "Java"
	subjectsSlice = append(subjectsSlice, "Go", "Database")
	// subjectsSlice[0] = "Java" // 슬라이싱
	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("================")
	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
}
