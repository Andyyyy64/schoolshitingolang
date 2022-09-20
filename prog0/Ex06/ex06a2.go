package main

import "fmt"

func main() {
	var input,i int
	real := 1
	fmt.Printf("階乗を計算したい数を入力してください :")
	fmt.Scan(&input)
	if input > 13 || input < 0 {
		fmt.Printf("計算できません")
	} else {
		for i = input; i > 0; i-- {
			real *= i;
		} 
		fmt.Printf("%d\n",real)
	}
}
