package main

import "fmt"

func main() {
	var i int
	n := [15]int{1, 7, 10, 18, 27, 31, 42, 49, 54, 50, 61, 67, 72, 88, 93}
	fmt.Printf("添字を入力してください :")
	fmt.Scan(&i)
	fmt.Printf("array[%d] = %d\n", i, n[i])
	if i >= 0 && i < 15 {
		if n[i]%3 == 0 {
			fmt.Printf("3の倍数です\n")
		} else {
			fmt.Printf("３の倍数ではありません\n")
		}
	}else{
        fmt.Printf("範囲外です\n")
    }
}
