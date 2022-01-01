package main

import (
	"fmt"
)

// struct tanımı alttaki gibi yapılır
type person struct {
	IsAlive        *bool
	IsAliveVal     bool
	Name           string
	MyCustomPerson *person // person tipli bir struct içindeki MyCustomPerson propertysi başka bir person tipli
	// pointer tutar
}

//structtaki bool pointer prop'una direkt olarak true false atanmaz çünkü adres ister.değerin adresi olmaz
func b(val bool) *bool { //primitive tip olan bool pointerı dönmek için 2 yol var
	//1
	//yx := val
	//return &yx
	//2
	x := new(bool)
	*x = val
	return x
	//pass by value - pass by reference func'a gonderilen pointer ise gosterdıgı degisken geriye donebilir
}
func main() {
	//var pI *int //int tipli pointer tanımı
	var I int = 4
	//pI = &I // değeri 4 olan int in ram adresini int pointer tipteki pIya atama
	Increment(&I)
	person2 := person{IsAliveVal: true, Name: "DoğancanKasap"}
	person1 := person{MyCustomPerson: &person2, IsAlive: b(person2.IsAliveVal), IsAliveVal: !person2.IsAliveVal, Name: "Dogan"}
	fmt.Println(*(person1).IsAlive)
	fmt.Println(person1.IsAliveVal)
	fmt.Println(person1.MyCustomPerson) // atanmayan değer zero value alır
	fmt.Println(person2.Name)
	person2.Name = "test" // structlar mutable dır. değiştirilebilir
	fmt.Println(person2.Name)
}

func Increment(pointerI *int) {
	*pointerI++ //dereferencing: pointer değişkeninin tuttuğu adresteki değere erişme/erişilen değeri arttırma
	fmt.Println(*pointerI)
}
