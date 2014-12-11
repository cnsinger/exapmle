package main

import "fmt"

func main() {
	messages := make(chan string)
	single := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	go func() {
		for {
			single <- true
			fmt.Println("waiting....")
			fmt.Println(<-messages)
			fmt.Println("success....")
			single <- true
			break
		}
		single <- true
	}()

	<-single
	select {
	case messages <- msg:
		fmt.Println("sent message success.", msg)
	default:
		fmt.Println("no message sent,again")
	}

	<-single
	select {
	case messages <- msg:
		fmt.Println("sent message success.", msg)
	default:
		fmt.Println("no message sent,again and again")
	}
	<-single
}
