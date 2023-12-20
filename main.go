package main

import "fmt"

func main() {

	channel := make(chan int)

	go setNumbers(channel)

	for v := range channel {
		fmt.Println(v)
	}
}

	func setNumbers(channel chan int){
		for i:= 0; i< 10; i++ {
			channel <- i
		}
		close(channel)
	}
