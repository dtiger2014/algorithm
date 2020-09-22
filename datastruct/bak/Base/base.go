package main

import(
	"fmt"
)

func main() {
	// array -------------------------
	var a1 = [10]int {1,2,3,4,5,6,7,8,9,10}
	var a2 = [10]int {1,2,3,4,5}
	var a3 = [...]int {1,2,3,4,5,6}
	var a4 = [10]int {2:4, 5:8}
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)	
	fmt.Println("len:", len(a3))

	var i, j, row, col, max int
	var a = [3][4]int {{1, 3, 7, 3}, {2, 3, 7 , 9}, {22, 3, 5, 10}}
	max = a[0][0]
	for i = 0; i <= 2; i ++ {
		for j = 0; j <= 3; j++ {
			if a[i][j] > max {
				max = a[i][j]
				row = i
				col = j
			}
		}
	}
	fmt.Println(a)
	fmt.Printf("max = %d, row = %d, col = %d\n", max, row, col)
	fmt.Println("=============================")

	// slice ------------------------------
	var s1 []int
	fmt.Printf("s1=%v, len=%v, cap=%v\n", s1, len(s1), cap(s1))

	s2 := []int{1,2,3}
	fmt.Printf("s2=%v, len=%v, cap=%v\n", s2, len(s2), cap(s2))
	for i := 0; i < 10; i++{
		s2 = append(s2, i)
		fmt.Printf("append s2=%v, len=%v, cap=%v\n", s2, len(s2), cap(s2))
	}
	for i:= 0; i < 5; i++{
		s2 = removeNoOrder(s2, i)
		fmt.Printf("remove s2=%v, len=%v, cap=%v\n", s2, len(s2), cap(s2))
	}	

	fmt.Println("===============================")

	// map ---------------------------------
	var m1 map[int]string
	fmt.Println(m1)

	m2 := map[int]string {0:"piao", 10:"fanfan"}
	fmt.Printf("map m2=%v, len=%v \n", m2, len(m2))

	m3 := make(map[int]string)
	m3[1], m3[2] = "Hello", "World"
	fmt.Printf("map m3=%v, len=%v \n", m3, len(m3))

	delete(m3, 1)
	fmt.Println(m3)

	fmt.Println("===============================")

	m := map[string]string{"key1":"val1"}  
    fmt.Println(m)  
    m["key2"] = "val2"  
    fmt.Println(m)  
    p := m
    fmt.Println(p)  
    delete(m, "key1")  
	fmt.Println(m) 
	fmt.Println(p) 

}

func removeFirst(s []int) []int {
	return s[1:] 
}

func remove(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

func removeNoOrder(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func removeLast(s []int) []int{
	return s[:len(s)-1]
}