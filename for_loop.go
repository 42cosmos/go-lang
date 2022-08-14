package main

import (
	"fmt"
	"time"
)

func read_word(ch chan string) {
	fmt.Println("Type a word, then hit Enter !")
	var word string
	fmt.Scanf("%s", &word)
	ch <- word
}

func print_char() {
	for {
		fmt.Printf(".")
		// 2초에 한 번씩 점을 찍어줌
		time.Sleep(2 * time.Second)
	}
}

func main() {
	defer fmt.Println("=====BYE . . .")

	go print_char()

	ch := make(chan string)
	go read_word(ch)
	// select == switch
	// 받은 메세지를 받아서 바로 출력하게끔 구현
	select {
	case word := <-ch:
		fmt.Println("\nReceived: ", word)
	}
}
