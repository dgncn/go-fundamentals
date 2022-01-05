package main

import "fmt"

func main() {
	bufChan := make(chan int, 3) // buffered chan tanımı. 3 buffer sizedır
	bufChan <- 1                 // channel FIFO şeklinde işler
	bufChan <- 2
	fmt.Println(<-bufChan)
	fmt.Println(<-bufChan)
	//fmt.Println(<- bufChan) buffer channel boş olduğu için deadlock olur
	// ya da 3 buffer sizelı channela 3ten fazla veri gönderimi deadlock oluşturur bloklanır
	// buffered channellar tam dolu veya tam boş durumlarda deadlock olusur gönderilen ve alınan listesi
	// buffer size içindeyse sorun olmaz

	fmt.Println("================= select ===================")
	//select çoklu channeldan veri alır veya gönderir

}
