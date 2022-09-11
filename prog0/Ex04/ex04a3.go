package main

import "fmt"

func main() {
	var score int

	fmt.Println("点数を入力してください ")
	fmt.Scan(&score)
	fmt.Printf("点数は %d です\n",score)
	
	if score >= 50 {
		fmt.Printf("合格")
	} else {
		fmt.Printf("不合格")
	}
}
