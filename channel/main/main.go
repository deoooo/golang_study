package main

import "fmt"

func main() {
	var intChan chan int = make(chan int, 10)
	intChan <- 10

	fmt.Println(1)
	fmt.Println(<-intChan)
	fmt.Println(2)
	fmt.Println(<-intChan)
	fmt.Println(3)
}
