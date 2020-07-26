package main

// 利用goroute判断数组内的数组个数L
func main() {
	myChan := make(chan int, 1000)
	for i := 0; i < 10000; i++ {
		myChan <- i
	}

	resChan := make(chan bool, 1)
}
