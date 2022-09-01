package main

import (
	"fmt"
	"log"

	"practice/queue"
	"practice/stack"
)

func main() {
	que := queue.New(stack.Stack{}, stack.Stack{})
	err := que.Push(4)
	if err != nil {
		log.Println(err.Error())
	}

	err = que.Push(2)
	if err != nil {
		log.Println(err.Error())
	}

	err = que.Push(6)
	if err != nil {
		log.Println(err.Error())
	}

	ele, err := que.Pop()
	if err != nil {
		fmt.Println("queue is empty right now...popped element will be", ele)
	} else {
		fmt.Println("popped element from queue is", ele)
	}
}
