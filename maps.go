package main

import "fmt"

type animal struct {
	FootCount int
	TypeName  string
}

func main() {
	/*
		map c#taki dictionary gibidir üç şekilde tanımlanır.
	*/
	//1
	var myMap map[string]int = make(map[string]int)
	myMap["first"] = 1
	myMap["second"] = 2
	//2
	secondMap := make(map[string]int)
	secondMap["1"] = 11
	fmt.Println(myMap, secondMap)

	/*
		virgülden sonraki değişken ile o key'e ait value var mı kontrolü yapılır
		eğer key varsa ok değişkeni true alır ve key değeri myValue'ya atanır.
		eğer keye karşı bir değer yoksa myValue int default olarak 0 alır
		eğer myValue değeri kullanımına ihtiyaç yok ise _ ile değiştirilir
		eğer ok değişkeni kullanılmıyorsa unused hatası döner. bu durum if koşuluyla kontrol edilir
	*/
	//_,ok := myMap["first5"]
	//myValue,ok := myMap["first5"]
	//fmt.Println(ok, myValue)
	//alttaki ornek: ok true ise dogru atama yapılmıstır ve terminale output verir
	if myValue, ok := myMap["first"]; ok == true {
		fmt.Println(myValue)
	}
	//3
	thirdMap := map[string]int{
		"first":  11,
		"second": 22,
	}
	fmt.Println(thirdMap)
	//mapten item silme
	delete(thirdMap, "second")
	fmt.Println(thirdMap)
	fmt.Println("-----------------------")

	dog := animal{
		FootCount: 2,
		TypeName:  "Karabaş",
	}
	fmt.Println(dog.TypeName)
}
