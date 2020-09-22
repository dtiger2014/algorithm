package functions

import(
	util "util"
	"time"
	"fmt"
)

func TestArrayQueue() {
	arrayQueue := util.ArrayQueue(10)

	for i := 0; i < 10; i++ {
		arrayQueue.Enqueue(i);
		arrayQueue.ToString()
		if i % 3 == 2 {
			arrayQueue.Dequeue()
			arrayQueue.ToString()
		}
	}
}

func ArrayQueueEfficiency(cnt int) {
	start := time.Now()
	// arrayqueue ---
	queue := util.ArrayQueue(10)
	
	for i := 0; i < cnt; i++ {
		queue.Enqueue(i)
	}

	for i := 0; i < cnt; i++ {
		queue.Dequeue()	
	}
	// -----------
	useTime := time.Since(start)
	fmt.Println("ArrayQueue Use time : ", useTime)
}