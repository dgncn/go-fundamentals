package main

import (
	"fmt"
)

//interface tanımı
type carInterface interface {
	accelerate()
	decelerate()
}

//carInterface'ini uygulayan tip tanımı
type car struct{}

//method tanımı. car değişkeni tarafından kullanılacak bir method
func (myCar car) accelerate() {
	fmt.Println("car is accelerating")
}

func (myCar car) decelerate() {
	fmt.Println("car is decelerating")
}

func main() {
	var carInt carInterface
	carInt = car{}
	carInt.accelerate()
	carInt.decelerate()

}
