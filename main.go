package main

import (
	"fmt"

	"practice/queue"
)

func main() {
	que := queue.New(nil, nil)
	que.Push(4)
	que.Push(2)
	fmt.Println(que.S1[len(que.S1)-1])
	ele := que.Pop()
	fmt.Println(ele)
}
