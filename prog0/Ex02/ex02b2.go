package main 

import "fmt"

func main(){
	var num,num2 int
	fmt.Println("2つ整数を入力してください")
	fmt.Scan(&num,&num2)
	fmt.Printf("和 = %d\n",num+num2)
	fmt.Printf("差 = %d\n",num - num2)
	fmt.Printf("積 = %d\n",num * num2)
	fmt.Printf("商 = %d、余り = %d\n",num/num2,num%num2)
}