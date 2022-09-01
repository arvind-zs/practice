package main

import (
	"fmt"

	"practice/queue"
	"practice/stack"
)

func main() {
	que := queue.New(stack.Stack{}, stack.Stack{})
	que.Push(4)
	que.Push(2)
	que.Push(6)
	ele := que.Pop()
	fmt.Println(ele)
}
