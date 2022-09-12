package main

import "fmt"

func main() {
	var n,r int
	count := 0

	fmt.Printf("初期値の入力 :")
	fmt.Scan(&n)
	for n > 0 {
		count ++
		r = n % 2
		n /= 2
		fmt.Printf("%d回目のループ 商 %d 余り %d\n",count,n,r)
	}

}