package main

import "fmt"

var x = 42

func main() {
	fmt.Println(x)
	foo()
	y := 17
	fmt.Printf(y)

}

func foo() {
	fmt.Println(x)
//	fmt.Printf(y)
//	y := 34
}
