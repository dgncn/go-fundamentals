package main

import "fmt"

//defer lifo
//defer functaki tüm işlevler tamamlandığında calısır. panic defer oncesindeyse defer calısmaz
//recover ile kurtarılan durum sonrasında *1 nolu defer çalışır
func main() {
	//defer fmt.Println("World")
	defer fmt.Println("World2") //*1
	fmt.Println("Hello")
	testpanic()
	fmt.Println("World")
}

//panic example
func testpanic() {

	defer func() {
		//recover error oluşan durumdan kurtarma gorevi gorur
		if err := recover(); err != nil {
			fmt.Println("devam")
		}
	}()
	// diğer dillerdeki catch gibi throw gibidir
	panic("set panic")

}
