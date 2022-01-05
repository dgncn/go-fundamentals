package main

import (
	"fmt"
)

func main() {

	//goroutine ile iletişim için channel kullanılır
	quitSignal := make(chan bool)
	go SayHiFromGoroutine(quitSignal) // goroutine için 1. yontem

	// her eş zamanlı çalıştırma aktivitesi goroutine olarak adlandırılır

	valueOfGoroutine := <-quitSignal
	if valueOfGoroutine == true {
		fmt.Println("hey2")
	} else {
		fmt.Println("hey1")
	}

	//thread işletim sistemi içerisindeki bir process'te çalışacak olan programın kodunu işleyen biridir
	//her program bir mantığa sahiptir ve bu mantığın işletilme sorumluluğu thread'e aittir.
	//goroutine bir fonksiyondur. programda yer alan diğer goroutinelerden bağımsız ve eş zamanlı çalışan bir fonk.
	//goroutine donanım bağımsız iken thread donanım bağlıdır.
	quitSignal2 := make(chan int)
	go func() {
		var myIntValue int = 4
		fmt.Println(myIntValue)
		quitSignal2 <- myIntValue
	}() // anonymous function goroutine tanımlamanın 2. yontemi
	<-quitSignal2

	//3
	// unbuffered channel
	/*
		ic := make(chan int)
		fmt.Println("test11")
		go printFor(ic)
		fmt.Println("test12")

		//channeldan alınan değerlerin iterasyonu için kullanılan for-range döngüsü.
		for icitem := range ic {
			fmt.Println("test1i")

			fmt.Println(icitem)
		}
		fmt.Println("test13")
	*/
	fmt.Println("=============== buffered channel ====================")
	icBuffered := make(chan int, 3)

	icBuffered <- 1
	icBuffered <- 2
	icBuffered <- 3

	close(icBuffered)
	for ix := range icBuffered {
		fmt.Println(ix)
	}

}

/*
func printFor(ic chan int) {
	i := 0
	for i < 4 {
		ic <- i
		i++
		time.Sleep(1 * time.Second)
	}
	close(ic)
}
*/

func SayHiFromGoroutine(quitSignal chan bool) {
	fmt.Println("hi from goroutine")
	quitSignal <- false
}
