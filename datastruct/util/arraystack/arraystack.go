package arraystack

import (
	"fmt"
	"golang-datastruct/util/array"
)

// ArrayStack : array stack
type ArrayStack struct {
	data *array.ArrayData
}

// New : 创建并初始化，并返回ArrayStack
func New() *ArrayStack {
	return &ArrayStack{
		data: array.New(),
	}
}

// GetCapacity : 获取Capacity
func (stack *ArrayStack) GetCapacity() int {
	return stack.data.GetCapacity()
}

// GetSize : 获取Size
func (stack *ArrayStack) GetSize() int {
	return stack.data.GetSize()
}

// IsEmpty : 判断是否为空
func (stack *ArrayStack) IsEmpty() bool {
	return stack.data.IsEmpty()
}

// Push : 目标元素 压入栈
func (stack *ArrayStack) Push(e interface{}) {
	stack.data.AddLast(e)
}

// Pop : 取出 栈顶元素
func (stack *ArrayStack) Pop() (interface{}, error) {
	return stack.data.RemoveLast()
}

// Peek : 获取 栈顶元素
func (stack *ArrayStack) Peek() (interface{}, error) {
	return stack.data.GetLast()
}

// String : 格式化输出
func (stack *ArrayStack) String() string {
	str := fmt.Sprintf("Stack:\n")
	str += "["
	for i := 0; i < stack.data.GetSize(); i++ {
		v, _ := stack.data.Get(i)
		str += fmt.Sprintf("%v", v)
		if i != stack.data.GetSize()-1 {
			str += ", "
		}
	}
	str += "] top"

	return str
}
