package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 150)
	}
}

func letters() {
	for l := 'a'; l < 'j'; l++ {
		fmt.Printf("%c ", l)
		time.Sleep(time.Millisecond * 230)
	}
}

func main() {
	go numbers()
	go letters()

	// These 3 seconds are a way of waiting for the execution of the letters method to finish.
	// It's nothing more than a hack, an appropriate way to solve this would be to use workgroups or channels.
	time.Sleep(3 * time.Second)
	fmt.Println("End of execution !")
}
