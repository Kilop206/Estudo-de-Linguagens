package main

import "fmt"

func main() {

	var num int = 1
	fmt.Println(num)

	num = num + 1
	fmt.Println(num)

	num = num - 2
	fmt.Println(num)

	num = num * 3
	fmt.Println(num)

	num = num / 4
	fmt.Println(num)

	num = num % 5
	fmt.Println(num)

	num++
	fmt.Println(num)

	num--
	fmt.Println(num)

	num += 1
	fmt.Println(num)

	num -= 1
	fmt.Println(num)

	num *= 1
	fmt.Println(num)

	num /= 1
	fmt.Println(num)

	num %= 1
	fmt.Println(num)

	num &= 1
	fmt.Println(num)

	num |= 1
	fmt.Println(num)

	num ^= 1
	fmt.Println(num)

	num <<= 1
	fmt.Println(num)

	num >>= 1
	fmt.Println(num)

	fmt.Println(num == 1)

	fmt.Println(num != 1)

	fmt.Println(num > 0)

	fmt.Println(num < 0)

	fmt.Println(num >= 0)

	fmt.Println(num <= 0)

	fmt.Println(num <= 0 && num > -1)

	fmt.Println(num <= 0 || num >= 1)

	fmt.Println(!(num == 1))

	fmt.Println(num & 1)

	fmt.Println(num | 1)

	fmt.Println(num ^ 1)

	fmt.Println(num << 1)

	fmt.Println(num >> 1)
}