package main

import "fmt"

func main() {
	var low,height,index float64
	fmt.Println("三角形の底辺の長さと高さをcmで入力してください :")
	fmt.Scan(&low,&height)
	index = (low * height) / 2
	fmt.Printf("三角形の面積は %f 平方cmです",index)
}
