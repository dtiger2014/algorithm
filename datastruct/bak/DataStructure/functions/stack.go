package functions

import(
	util "util"
	"fmt"
	"time"
)

func TestStack() {
	stack := util.ArrayStack(10)

	for i := 0; i < 5; i++ {
		stack.Push(i)
		stack.ToString()
	}

	stack.Pop()
	stack.ToString()
}

func ArrayStackEfficiency(cnt int) {
	start := time.Now()
	// loopqueue ---
	stack := util.ArrayStack(10)
	
	for i := 0; i < cnt; i++ {
		stack.Push(i)
	}

	for i := 0; i < cnt; i++ {
		stack.Pop()	
	}
	// -----------
	useTime := time.Since(start)
	fmt.Println("ArrayStack Use time : ", useTime)
}

func TestStack2(s string) bool {
	stack := util.ArrayStack(20)

	for _, v := range s {
		c := string(v)
		if c == "(" || c == "[" || c =="{" {
			stack.Push(c)
		} else {
			if stack.IsEmpty() {
				fmt.Println("Mismatch! string=", s)
				return false
			}

			topC := stack.Pop()
			if c == ")" && topC != "(" {
				fmt.Println("Mismatch! string=", s)
				return false
			}
			if c == "]" && topC != "[" {
				fmt.Println("Mismatch! string=", s)
				return false
			}
			if c == "}" && topC != "{" {
				fmt.Println("Mismatch! string=", s)
				return false
			}
		}
	}

	if stack.IsEmpty() {
		fmt.Println("Match! string=", s)
		return true
	}

	fmt.Println("Mismatch! string=", s)
	return false
}