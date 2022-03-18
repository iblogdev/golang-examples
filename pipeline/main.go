package main

import "fmt"

func main() {
	inputs := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	outputStream := stage3(stage2(stage1(inputs...)))

	for output := range outputStream {
		fmt.Println(output)
	}
}

func stage1(inputStream ...int) chan int {
	outputStream := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			outputStream <- i
		}
		close(outputStream)
	}()
	return outputStream
}

func stage2(inputStream <-chan int) chan int {
	outputStream := make(chan int)
	go func() {
		for input := range inputStream {
			outputStream <- (input + 1)
		}
		close(outputStream)
	}()
	return outputStream
}

func stage3(inputStream <-chan int) chan int {
	outputStream := make(chan int)
	go func() {
		for input := range inputStream {
			outputStream <- (input * 2)
		}
		close(outputStream)
	}()
	return outputStream
}
