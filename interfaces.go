package main

import "fmt"

func main() {

	//interface tanımı
	var myInterface MyInterface = &MyInterfaceImplement{Name: "Ornek struct"}
	fmt.Println(myInterface.MyMethod())
}

type MyInterface interface {
	MyMethod() string
}

type MyInterfaceImplement struct {
	Name string
}

func (imp *MyInterfaceImplement) MyMethod() string {
	return imp.Name // pointer indirection (*imp).Name
}
