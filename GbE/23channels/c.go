package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		for {
			// messages <- fmt.Sprintf("ping #%d", i)
			var input string
			fmt.Scanln(&input)
			messages <- input
		}
	}()

	for {
		r := <-messages
		if r == "0" {
			return
		}
		fmt.Println(r)
	}
}
