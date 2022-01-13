package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("İşlem tipini giriniz:Toplama (T) Çıkarma(C)... ")
	processType, _ := reader.ReadString('\n')
	processType = strings.TrimSpace(processType)

	fmt.Print("1.sayıyı giriniz: ")
	s, _ := reader.ReadString('\n')
	num1, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Print("2.sayıyı giriniz: ")
	s2, _ := reader.ReadString('\n')
	num2, err := strconv.Atoi(strings.TrimSpace(s2))
	if err != nil {
		fmt.Println(err)
	}
	switch processType {
	case "T":
		fmt.Printf("%v", num1+num2)
	case "C":
		fmt.Printf("%v", num1-num2)
	}
}
