package functions

import(
	util "util"
	// "strconv"
	// "fmt"
	// "time"
)

func TestLinkedList() {
	linkedList := util.LinkedList()

	for i := 0; i < 5; i++ {
		linkedList.AddLast(i)
		linkedList.ToString()
	}
	linkedList.Add(2, 666)
	if linkedList.Contains(666) {
		linkedList.Set(2, 333)
	}
	linkedList.ToString()

	linkedList.Remove(2)
	linkedList.ToString()

	linkedList.RemoveFirst()
	linkedList.ToString()

	linkedList.RemoveLast()
	linkedList.ToString()

	linkedList.Add(3, 555)
	linkedList.ToString()

	linkedList.RemoveElement(3)
	linkedList.ToString()
}