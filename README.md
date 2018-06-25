# coordinator

A wrapper around WaitGroup so you don't have to increment the delta.

## Usage

```golang
coordinator.Run(func () {
  <runs in a goroutine>
})
coordinator.Wait()
```

## Example

```golang
package main

import (
	"fmt"

	co "github.com/evanisnor/coordinator"
)

func main() {
	messages := make(chan string)

	emit(messages, "ping", 2)
	emit(messages, "pong", 2)
	go print(messages)

	co.Wait()
}

func emit(messages chan<- string, word string, count int) {
	co.Run(func() {
		id := 0
		for id < count {
			messages <- fmt.Sprintf("%s %d", word, id)
			id++
		}
	})
}

func print(messages <-chan string) {
	for val := range messages {
		fmt.Printf("%v\n", val)
	}
}

```