package main

import "fmt"

func main() {
	var id,han int
	fmt.Println("学籍番号の下三桁を入力してください(例: 001)")
	fmt.Scan(&id)
	han = id % 3
	switch han {
	case 1:
		fmt.Printf("この人は1班です\n")
	case 2:
		fmt.Printf("この人は2班です\n")
	case 0:
		fmt.Printf("この人は3班です\n")
	}
}
