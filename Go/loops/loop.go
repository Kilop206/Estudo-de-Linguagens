package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {

		if i == 1 {
			break
		}

		if (i == 2) {
			continue
		}

		fmt.Println(i)
	}

	array := [...]int{1,2,3,4,5} 
	for i := range array {
		fmt.Println(i)
	}
}