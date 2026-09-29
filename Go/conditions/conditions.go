package main

import "fmt"

func main() {
	var num int = 1

	if num == 1 {
		fmt.Println(1)
	} else if num == 2 {
		fmt.Println(2)
	} else {
		fmt.Println(3)
	}

	switch num {
		case 1 : 
				fmt.Println(1)
		case 2 : 
			fmt.Println(2)
		default :
			fmt.Println(3)
	}
}