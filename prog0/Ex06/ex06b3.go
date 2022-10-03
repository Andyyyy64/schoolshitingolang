package main

import "fmt"

func main() {
	var n int
	fmt.Printf("1から9までの整数を入力してください :")
	fmt.Scan(&n)
	if n <= 0 || n >= 10 {
		fmt.Printf("入力する数は1から9までです")
	} else {
		for i := 1; i < n; i++ {
			for j := 0; j < i; j++ {
				fmt.Printf("%d", i)
			}
			fmt.Printf("\n")
		}
		for i := n; i > 0; i-- {
			for j := 0; j < i; j++ {
				fmt.Printf("%d",i)
			}
			fmt.Printf("\n")
		}
	}

}
