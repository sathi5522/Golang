package main

import "fmt"

var q int = 90

func main() {
	var a int = 10
	fmt.Println("a :", a)

	var s string = "satheesh"
	//fmt.Println("s : ", s)
	_ = s

	s = "kumar"
	b := 20
	str := "nelavalli"
	fmt.Println("b:", b, "str: ", str)

	//multiple declarations

	car, cost := "Baleno", 900000
	fmt.Println(car, cost)
	car, year := "Baleno", 2019
	_ = year

	var opened = false
	opened, file := true, "a.txt"
	fmt.Println(opened, file)

	var (
		name    string
		id      int
		company string
	)
	fmt.Println(name, id, company)

	var x, y, z int
	fmt.Println(x, y, z)

	var h, k int
	h, k = 3, 4
	h, k = k, h //swapping variables
	fmt.Println(h, k)

	sum := 5 + 2.3
	fmt.Println(sum)
}
