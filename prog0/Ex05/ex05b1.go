package main

import "fmt"

func main() {
	var score, total int
	var ave float64
	student := 0
	pas := 0
loop:
	for true {
		fmt.Printf("得点を入力してください :")
		fmt.Scan(&score)
		if score < 0 {
			break loop
		}
		student++
		total += score
		if score >= 50 {
			pas++
		}
	}
	ave = float64(total / student)
	fmt.Printf("受験者数 : %d  平均点 : %f\n", student, ave)
	fmt.Printf("合格者 %d 不合格者 %d\n", pas, student-pas)
}
