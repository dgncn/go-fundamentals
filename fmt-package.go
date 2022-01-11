package main

import "fmt"

type Myfloat float32

type MyStruct struct {
	id string
}

func main() {
	var x string = "hey2"
	fl := Myfloat(288)
	mys := &MyStruct{id: x + "sss"}
	//T: Tip
	fmt.Printf("%v %v %T %t %b %q %x %X \n", x, fl, mys.id, "1" == "2", 16, "test", "1aa", "2cc")

	s2 := struct {
		i  int
		co complex64
	}{i: 22, co: complex(3, 4)}
	fmt.Printf("%+v - %+v\n", s2, s2.i)
	fmt.Printf("%+v %+v\n", s2.co, real(s2.co)) //karmaşık sayı. 3+4i
	fmt.Printf("%v", s2)

	s := ""
	i := 0
	fmt.Sscanf("12345", "%3s%d", &s, &i) //3 karakteri s stringine geri kalanı %d / int tipte iye atar
	fmt.Printf(" i: %v itype: %T s: %v stype: %T", i, i, s, s)
}
