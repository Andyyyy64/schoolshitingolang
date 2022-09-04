package main

import "fmt"

func main(){
	var number int
	fmt.Println("整数値を入力してください")
	fmt.Scan(&number)
	fmt.Printf("入力された値は%dで、これを5で割った余りは%dです\n",number,number%5)
}