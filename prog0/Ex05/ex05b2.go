package main

import "fmt"

func main() {
	var input int
	count := 0

	for true {
		fmt.Printf("偶数は%d回入力されています 正の整数を入力してください :", count)
		fmt.Scan(&input)
		if input%2 == 0 {
			count++
			if count == 2 {
				goto fini
			}
		}
	}
fini:
	fmt.Printf("2回偶数が入力されました。プログラムを終了します \n")
}
