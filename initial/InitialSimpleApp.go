package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	//fist thing that get executed in program
	fmt.Println("hey Ronak")

	var intNumber = 13
	//var floatNumber float64 = 7

	fmt.Println("memory location :- ", &intNumber)

	//we cannot write string in next line with "", if needed then do concatenate(+)
	//but we can do that using ``
	var String string = "hey boy"

	fmt.Println(String)

	//get length of String
	fmt.Println(utf8.RuneCountInString(String))

	//var mybook bool = false

	userName := "Ronak"
	fmt.Println(userName)

	//constant are like private static final in java
	const pi float64 = 3.14

	//call the functions from main function
	printMe("Hey user")

	//fmt.Println(intDivision(30, 6))

	//store two return type values in 2 variables

	var result, remainder, err = intDivision(30, 0)

	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The value of remainder is %v", remainder)
	} else {
		//we can use printf function
		fmt.Printf("division is %v with remainder %v", result, remainder)
	}

	//array
	var arr = [3]int{1, 2, 3}
	fmt.Println("\n", arr)

	/*for i := 0; i < len(arr); i++ {
		fmt.Println("Value is : ", arr[i])
	}*/

	for value := range arr {
		fmt.Println("array: ", arr[value])
	}

	//slices are just wrapper of array under the hood
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	//get length of slice
	fmt.Println("Length is : ", len(slice))

	//append to the slice
	fmt.Println(append(slice, 20, 30, 40, 50))

	fmt.Println("Slicing is: ", slice[1:3])

	//map
	var myMap = make(map[int]string)
	myMap[1] = "Ronak"
	myMap[2] = "Ronak1"
	myMap[3] = "Ronak2"
	myMap[4] = "Ronak3"
	myMap[5] = "Ronak4"

	fmt.Printf("\t%v", myMap)

	for key, element := range myMap {
		fmt.Println(key, " : ", element)
	}
}

// functions
func printMe(printValue string) {
	fmt.Println(printValue)
}

//we can define return type like this after the parameters
/*func intDivision(numerator int, denominator int) int {
	return numerator / denominator
}*/

// can define multiple return type
func intDivision(numerator int, denominator int) (int, int, error) {

	var err error
	if denominator == 0 {
		err = errors.New("Can't divide by zero")
		return 0, 0, err
	}

	return numerator / denominator, numerator % denominator, err
}
