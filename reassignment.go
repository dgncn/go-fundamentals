package main

import "fmt"

func main() {
	//reassignment: := ile tanımlı bir degiskene birden fazla kez atama yapmak.
	total, err := aggregateNonZeroValues(1, 0)
	//err variable declared but if err is nil, it can be assigned value to err variable in beyond(line 12)
	//this is legal.
	if err != nil {
		//panic(err)
	}
	fmt.Println(total)
	fmt.Println(err)
	total2, err := aggregateNonZeroValues(0, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(total2)
}

func aggregateNonZeroValues(number1, number2 int) (int, *error) {
	if number1 == 0 || number2 == 0 {
		return 0, new(error)
	} else {
		return number1 + number2, nil
	}
}
