package main

import (
	"fmt"
)

//interface tanımı
type carInterface interface {
	accelerate()
	decelerate()
	increaseSpeed(speed int) *car
	getSpeedValue() int
	get()
}

//carInterface'ini uygulayan tip tanımı
type car struct {
	_speed int
}

//method tanımı. car değişkeni tarafından kullanılacak bir method
func (myCar car) accelerate() {
	fmt.Println("car is accelerating")
}

func (myCar car) decelerate() {
	fmt.Println("car is decelerating")
}

//car nesnesinin referansı üstünde field set yapar ve aldığı nesneyi döner
func (myCar *car) increaseSpeed(speedVal int) *car {
	(*myCar)._speed = speedVal
	return myCar
}

func (myCar *car) getSpeedValue() int {
	return (*myCar)._speed
}

func (myCar *car) get() {
	fmt.Println(myCar)
}

func main() {
	var carInt carInterface
	carInt = &car{}
	carInt.accelerate()
	carInt.decelerate()
	carInt.increaseSpeed(150)
	fmt.Println(carInt.getSpeedValue())
	carInt.get()
}
