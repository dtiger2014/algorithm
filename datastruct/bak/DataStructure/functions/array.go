package functions

import(
	util "util"
	
	"strconv"
	"fmt"
	"time"
)

func MyArrayTest() {
	start := time.Now()
	// myarray ---
	array := util.MyArray(10)
	cnt := 10000
	for i := 0; i < cnt; i++ {
		array.AddLast(i)
		// array.AddFirst(i)
	}

	for i := 0; i < cnt; i++ {
		array.RemoveLast()
		// array.RemoveFirst()
	}
	array.ToString()
	// -----------
	useTime := time.Since(start)
	fmt.Println("Use time : ", useTime)
}

func MyArrayInt() {
	array := util.MyArray(2)

	for i := 0; i < 10; i++ {
		array.AddLast(i)
	}
	array.ToString()

	array.Add(1, 100)
	array.ToString()

	array.AddFirst(-1)
	array.ToString()

	array.Remove(2)
	array.ToString()

	array.RemoveElement(4)
	array.ToString()

	array.RemoveFirst()
	array.ToString()
}

func MyArrayString() {
	array := util.MyArray(20)

	for i := 0; i < 10; i++ {
		array.AddLast(strconv.Itoa(i))
	}
	array.ToString()

	array.Add(1, "100")
	array.ToString()

	array.AddFirst("-1")
	array.ToString()

	array.Remove(2)
	array.ToString()

	array.RemoveElement("4")
	array.ToString()

	array.RemoveFirst()
	array.ToString()
}

type student struct{
	name string
	score int
}

func MyArrayStudent() {
	array := util.MyArray(20)
	array.AddLast(student{name:"student1",score:100})
	array.AddLast(student{name:"student2",score:80})
	array.AddLast(student{name:"student3",score:70})
	array.ToString()
}