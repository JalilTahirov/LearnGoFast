package main

import "fmt"

func main() {

	/*
		var number float32

		fmt.Print("USD: ")

		_, err := fmt.Scanf("%f", &number)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			return
		}

		//result1:= 3.4
		var result2 float64 = 3.4
		//result = number * 1.7
		fmt.Println("AZN:", result2)
	*/

	var arrInput = []int{1, 2, 3}

	//var restult = addFunction(5)
	var restult = addFunction2(arrInput)

	fmt.Println(restult)
}

func myMessage() {
	fmt.Println("I just got executed!")
}

func addFunction(value int) []int {
	var array = []int{}
	fmt.Println(value)
	array = append(array, value)
	return array
}

func addFunction2(values []int) []int {
	var array = []int{}

	for i := 0; i < len(values); i++ {
		fmt.Println(values[i])
		array = append(array, values[i])
	}
	//

	//array = append(array, values)
	return array
}
