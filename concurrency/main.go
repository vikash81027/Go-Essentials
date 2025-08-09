package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)
} 

func main() {
	// dones := make([]chan bool, 4)

	// for idx := range dones {
	// 	dones[idx] = make(chan bool)
	// }

	// go greet("Hello 1", dones[0]);
	// go greet("Hello 2", dones[1]);
	// go slowGreet("slow Hello 1", dones[2])
	// go greet("Hello 3", dones[3]);

	// for idx := range 4 {
	// 	<- dones[idx]
	// }

	done := make(chan bool)

	go greet("Hello 1", done);
	go greet("Hello 2", done);
	go slowGreet("slow Hello 1", done)
	go greet("Hello 3", done);

	for doneChan := range done {
		fmt.Println(doneChan)
	}
	
}