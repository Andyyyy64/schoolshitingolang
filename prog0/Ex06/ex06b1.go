package main

import "fmt"

func main() {
	i := 0
	n := 0
	sum := 0
	target := 47
	hit := 0	 
	for i = 1; i <= 10 ; i++ {
		fmt.Printf("1から9までの数を入力してください。%2d回目: ", i)
		fmt.Scan(&n)
		/* 指定された範囲外の数が入力されたら、入力をやり直す（回数のカウントは増える） */
		if n >= 10 || n <= 0 {
			fmt.Printf("その数字は範囲外です\n")
			continue
		}

		sum += n
		if sum == target {
			hit = 1
			break
		}
		if sum > target {
			break
		}
	}
	if hit == 1 {
		fmt.Printf("%d回目で当たりました\n\n",i)
	 } else if  sum < target{
		fmt.Printf("残念ですが、入力された数の合計(%d)が設定未満でした\n\n",sum)
	 } else {
		fmt.Printf("残念ですが、入力された数の合計(%d)が設定値を超えました。\n\n",sum)
	 }

}
