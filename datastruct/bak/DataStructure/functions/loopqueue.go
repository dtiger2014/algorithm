package functions

import(
	util "util"
	"fmt"
	"time"
	"strconv"
)

func TestLoopQueue() {
	loopQueue := util.LoopQueue(10)

	for i := 0; i < 10; i++ {
		loopQueue.Enqueue(i)
		loopQueue.ToString()

		if i % 3 == 2 {
			loopQueue.Dequeue()
			loopQueue.ToString()
		}
	}
}

func TestLoopQueue2() {
	loopQueue := util.LoopQueue(10)

	for i := 0; i < 3; i++ {
		// loopQueue.Enqueue(i)
		str := "s" + strconv.Itoa(i)
		loopQueue.Enqueue(str)
		loopQueue.ToString()
	}
}

func LoopQueueEfficiency(cnt int) {
	start := time.Now()
	// loopqueue ---
	queue := util.LoopQueue(10)
	
	for i := 0; i < cnt; i++ {
		queue.Enqueue(i)
	}

	for i := 0; i < cnt; i++ {
		queue.Dequeue()	
	}
	// -----------
	useTime := time.Since(start)
	fmt.Println("LoopQueue Use time : ", useTime)
}