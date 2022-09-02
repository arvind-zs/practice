package main

import (
	"fmt"
	"log"

	"practice/queue"
	"practice/stack"
)

func main() {
	que := queue.New(stack.Stack{}, stack.Stack{})

	for i := 0; i < 101; i++ {
		err := que.Push(4)
		if err != nil {
			log.Println(err.Error())
		}
	}

	ele, err := que.Pop()
	if err != nil {
		fmt.Println("queue is empty right now...popped element will be", ele)
	} else {
		fmt.Println("popped element from queue is", ele)
	}
}
