package main

import (
	"fmt"
	// "learn_go/greet"
)

type Student struct{
	name string
	age int
}

func main(){
	fmt.Println("Hello From Main")
	for i :=0; i<=10000;i++{
		fmt.Println(i)
	}
}