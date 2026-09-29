package main

import "fmt"

func main() {
	var array = [1]int{1}

	fmt.Println(array)

	newArray := [7]int{1,2,3,4,5,6,7}

	fmt.Println(newArray)

	var otherArray = [...]int{1,2,3}

	fmt.Println(otherArray)
}