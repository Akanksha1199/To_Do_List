package main

import "fmt"

// // implimenting struct
// type Cat struct{}
// type dog struct{}

// // implementing functions is actually a method
// func (c Cat) Say() string {
// 	return "meow"
// }
// func (d dog) Say() string {
// 	return "woof"
// }

type number struct {
	num1 int
	num2 int
}
type myInterface interface {
	add(a, b int) int
}

func main() {
	fmt.Println("Interface Practice")

	// //calling a struct
	// c := Cat{}
	// d := dog{}
	// fmt.Println("Cat says ", c.Say())
	// fmt.Println("Dog says ", d.Say())

	//a := 10
	//b := 2
	//i := number{}
	//fmt.Println(i)
	var v myInterface = number{10, 20}
	fmt.Println(v.add(1, 2))
}
func (n number) add(a, b int) int {
	return (a + b*n.num1 + n.num2)
}
