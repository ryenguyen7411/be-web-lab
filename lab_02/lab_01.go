package main

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	fmt.Printf("Hello, %s dep trai vip pro!\n", name)
}

func main() {
	go sayHello("Rye Nguyen")
	go sayHello("Tam Le")
	go sayHello("Minh Tran")
	go sayHello("Thanh Nguyen")
	go sayHello("Loi Ly")
	go sayHello("Dung Nguyen")

	time.Sleep(1 * time.Second) // waiting for above process 1 seconds
	fmt.Println("Main finished")
}
