package functions

import(
	util "util"
	// "strconv"
	"fmt"
	"time"
)

func TestLinkedListQueue() {
	queue := util.LinkedListQueue()
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		queue.ToString()	

		if i % 3 == 2 {
			queue.Dequeue()
			queue.ToString()
		}
	}	
}

func LinkedListQueueEfficiency(cnt int) {
	start := time.Now()
	// arrayqueue ---
	queue := util.LinkedListQueue()
	
	for i := 0; i < cnt; i++ {
		queue.Enqueue(i)
	}

	for i := 0; i < cnt; i++ {
		queue.Dequeue()	
	}
	// -----------
	useTime := time.Since(start)
	fmt.Println("LinkedListQueue Use time : ", useTime)
}
