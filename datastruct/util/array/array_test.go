package array

import (
	"testing"
)

// func TestMyArray(t *testing.T) {
// 	tests := []struct {
// 		a   int
// 		ans int
// 	}{
// 		{7, 13},
// 		// {3, 4},
// 		{5, 5},
// 		// {13, 14},
// 	}

// 	for _, v := range tests {
// 		answer := Fib(v.a)
// 		if answer != v.ans {
// 			t.Errorf("Fib(%d) = %d; answer %d", v.a, v.ans, answer)
// 			// t.Fail()	//测试失败，测试继续，也就是之后的代码依然会执行
// 			// t.FailNow()	//测试失败，测试中断
// 			// t.SkipNow() //跳过测试，测试中断
// 			// t.Log() 	//输出信息
// 			// t.Logf() // 格式化输出信息

// 			// * $ go test -v . 打印信息

// 			// t.Error: 相当于 Log + Fail
// 			// t.Errorf: 相当于 Logf + Fail

// 			// t.Fatal: 相当于 Log + FailNow
// 			// t.Fatalf: 相当于 Logf + FailNow

// 			t.Log("test log :", 1, " number")
// 		} else {
// 			t.Logf("test success %d %d", v.a, v.ans)
// 		}
// 	}

// }

var pairs = []struct {
	k string
	v string
}{
	{"polaris", "徐新华"},
	{"studygolang", "Go语言中文网"},
	{"stdlib", "Go语言标准库"},
	{"polaris1", "徐新华1"},
	{"studygolang1", "Go语言中文网1"},
	{"stdlib1", "Go语言标准库1"},
	{"polaris2", "徐新华2"},
	{"studygolang2", "Go语言中文网2"},
	{"stdlib2", "Go语言标准库2"},
	{"polaris3", "徐新华3"},
	{"studygolang3", "Go语言中文网3"},
	{"stdlib3", "Go语言标准库3"},
	{"polaris4", "徐新华4"},
	{"studygolang4", "Go语言中文网4"},
	{"stdlib4", "Go语言标准库4"},
}

func TestWriteToMap(t *testing.T) {
	t.Parallel()
	for _, tt := range pairs {
		WriteToMap(tt.k, tt.v)
	}
}

func TestReadFromMap(t *testing.T) {
	t.Parallel()
	for _, tt := range pairs {
		actual := ReadFromMap(tt.k)
		if actual != tt.v {
			t.Errorf("the value of key(%s) is %s, expected: %s", tt.k, actual, tt.v)
		}
	}
}
