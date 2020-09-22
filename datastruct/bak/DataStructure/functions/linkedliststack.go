package functions

import(
	util "util"
	// "strconv"
	"fmt"
	"time"
)

func TestLinkedListStack() {
	stack := util.LinkedListStack()
	for i := 0; i < 5; i++ {
		stack.Push(i)
		stack.ToString()	
	}

	stack.Pop()
	stack.ToString()
}

func LinkedListStackEfficiency(cnt int) {
	start := time.Now()
	// loopqueue ---
	queue := util.LinkedListStack()
	
	for i := 0; i < cnt; i++ {
		queue.Push(i)
	}

	for i := 0; i < cnt; i++ {
		queue.Pop()	
	}
	// -----------
	useTime := time.Since(start)
	fmt.Println("LinkedListStack Use time : ", useTime)
}