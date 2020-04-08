package main

import (
	"fmt"
)

type myStr []byte

func (s myStr) Len() int           { return len(s) }
func (s myStr) Less(i, j int) bool { return s[i] < s[j] }
func (s myStr) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	arr := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println(isValidSudoku(arr))
}

func isValidSudoku(board [][]byte) bool {
	memo := make(map[string]int)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				continue
			}
			kHori := fmt.Sprintf("%v_h_%v", board[i][j], i)
			kVert := fmt.Sprintf("%v_v_%v", board[i][j], j)
			kBloc := fmt.Sprintf("%v_b_%v%v", board[i][j], i/3, j/3)

			if _, ok := memo[kHori]; ok {
				return false
			}
			if _, ok := memo[kVert]; ok {
				return false
			}
			if _, ok := memo[kBloc]; ok {
				return false
			}
			memo[kHori] = 1
			memo[kVert] = 1
			memo[kBloc] = 1
		}
	}
	return true
}
