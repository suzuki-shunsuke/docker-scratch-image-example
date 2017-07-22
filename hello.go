package main

import "fmt"
import "time"

func main() {
	for {
		fmt.Println("Hello")
		time.Sleep(1 * time.Second)
	}
}
