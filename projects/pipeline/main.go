package main

import "fmt"

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)
	go func() {
		inputStream <- "a"
		inputStream <- "a"
		inputStream <- "b"
		inputStream <- "b"
		inputStream <- "c"
		close(inputStream)
	}()
	for x := range outputStream {
		fmt.Print(x)
	}
	fmt.Print("\n")
}

func removeDuplicates(inputStream, outputStream chan string) {
	var Value string
	for v := range inputStream {
		if Value != v {
			outputStream <- v
			Value = v
		}
	}
	close(outputStream)
}
