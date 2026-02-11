package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	
	var intNum int16 = 32767
	intNum += 1
	fmt.Print(intNum)
	fmt.Println("The maximum value of int16 is:", intNum, "\n")

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 5
	result := floatNum32 + float32(intNum32)
	fmt.Println("The result of adding float32 and int32 is:", result)
	
	var intNum1 int = 10
	var intNum2 int = 7
	fmt.Println(intNum1/intNum2)
}
