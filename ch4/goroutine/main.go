package main

import (
	"fmt"
	"time"
)

func main() {
	go printNumbers()
	go printLetters()

	time.Sleep(3000 * time.Millisecond)
}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(250 * time.Millisecond)
	}
}

func printLetters() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c ", i)
		time.Sleep(400 * time.Millisecond)
	}
}
