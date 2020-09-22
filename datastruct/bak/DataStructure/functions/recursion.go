package functions

import(
	"fmt"
	util "util"
)

func TestSum() {
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("sum is", util.Sum(arr))
}
