package main

import (
	"fmt"
	"time"
)

func sayROutine(s string, done chan bool, fala chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		texto := "dados vindo de " + s
		fala <- texto
		fmt.Println(s)
	}
	done <- true
}

func sayROutineChan(s string, done chan bool, fala chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		resposta := <-fala
		fmt.Println("dado impresso em ", s, resposta)
		fmt.Println(s)
	}
	done <- true
}
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

}

func main() {
	done1 := make(chan bool)
	done2 := make(chan bool)
	fala := make(chan string)

	defer fmt.Println("Cabou")
	go sayROutine("hello", done1, fala)
	go sayROutineChan("world", done2, fala)
	<-done1
	<-done2
}
