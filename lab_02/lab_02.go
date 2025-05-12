package main

import (
	"fmt"
	"sync"
)

func sayHello(name string) {
	fmt.Printf("Hello, %s dep trai vip pro!\n", name)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		sayHello("Rye Nguyen")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sayHello("Tam Le")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sayHello("Minh Tran")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sayHello("Thanh Nguyen")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sayHello("Loi Ly")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sayHello("Dung Nguyen")
	}()

	wg.Wait()
	fmt.Println("Main finished")
}
