package main

import "fmt"

func main() {
	var temp int
	var hum float64
	fmt.Printf("温度と湿度を入力してください ")
	fmt.Scan(&temp, &hum)
	if temp < 20 {
		fmt.Printf("寒い")
	} else if hum >= 60 {
		fmt.Printf("蒸し暑い")
	} else if temp >= 28 {
		fmt.Printf("暑い")
	} else {
		fmt.Printf("快適")
	}
}
