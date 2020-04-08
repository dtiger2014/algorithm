package hash

import (
	"sort"
)

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	memo := make(map[int]int, len(nums))
	for _, v := range nums {
		if _, ok := memo[v]; ok {
			return true
		}
		memo[v] = 1
	}
	return false
}

func singleNumber(nums []int) int {
	memo := make(map[int]int, len(nums))
	for _, v := range nums {
		if _, ok := memo[v]; ok {
			memo[v]++
		} else {
			memo[v] = 1
		}
	}
	for i, v := range memo {
		if v == 1 {
			return i
		}
	}
	return -1
}

// 给定两个数组，编写一个函数来计算它们的交集。
// 输入: nums1 = [1,2,2,1], nums2 = [2,2]
// 输出: [2]
// 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// 输出: [9,4]
func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	memo := make(map[int]int)
	for _, v := range nums1 {
		if _, ok := memo[v]; !ok {
			memo[v] = 1
		}
	}

	for _, v := range nums2 {
		if _, ok := memo[v]; ok {
			memo[v] = 2
		}
	}

	ret := []int{}

	for num, cnt := range memo {
		if cnt == 2 {
			ret = append(ret, num)
		}
	}

	return ret
}

/*
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。如果 可以变为  1，那么这个数就是快乐数。

如果 n 是快乐数就返回 True ；不是，则返回 False 。
*/
/*
入：19
输出：true
解释：
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
*/
func isHappy(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	memo := make(map[int]int)
	memo[1] = 1
	tar, calNum := n, 0
	for tar != 1 {
		calNum = calHappyNum(tar)
		if v, ok := memo[calNum]; ok {
			return v == 1
		}
		memo[tar] = calNum
		tar = calNum
	}

	return true
}
func calHappyNum(num int) int {
	res, tar := 0, num
	for tar > 9 {
		res += (tar % 10) * (tar % 10)
		tar = tar / 10
	}
	res += tar * tar

	return res
}

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那
 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/
func twoSum(nums []int, target int) []int {
	if len(nums) <= 1 {
		return nil
	}

	memo := make(map[int]int)
	for i, v := range nums {
		if key, ok := memo[target-v]; ok {
			return []int{key, i}
		}
		memo[v] = i
	}
	return nil
}

/*
示例 1:

输入:
["Shogun", "Tapioca Express", "Burger King", "KFC"]
["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]
输出: ["Shogun"]
解释: 他们唯一共同喜爱的餐厅是“Shogun”。
示例 2:

输入:
["Shogun", "Tapioca Express", "Burger King", "KFC"]
["KFC", "Shogun", "Burger King"]
输出: ["Shogun"]
解释: 他们共同喜爱且具有最小索引和的餐厅是“Shogun”，它有最小的索引和1(0+1)。
*/
func findRestaurant(list1 []string, list2 []string) []string {
	memo := make(map[string]int)

	for i, str := range list1 {
		memo[str] = i
	}

	minIdxSum := int(^uint((0)) >> 1)
	for l2Idx, l2Str := range list2 {
		if l1Idx, ok := memo[l2Str]; ok {
			if l1Idx+l2Idx < minIdxSum {
				minIdxSum = l1Idx + l2Idx
			}
		}
	}

	res := []string{}
	for l2Idx, l2Str := range list2 {
		if l1Idx, ok := memo[l2Str]; ok && l1Idx+l2Idx == minIdxSum {
			res = append(res, l2Str)
		}
	}

	return res
}

/*
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:

s = "leetcode"
返回 0.

s = "loveleetcode",
返回 2.


注意事项：您可以假定该字符串只包含小写字母。
*/
func firstUniqChar(s string) int {
	memo := make([]int, 26)
	for _, c := range s {
		memo[c-'a']++
	}

	for i, c := range s {
		if memo[c-'a'] == 1 {
			return i
		}
	}
	return -1
}

/*
给定两个数组，编写一个函数来计算它们的交集。

示例 1:

输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2,2]
示例 2:

输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [4,9]
说明：

输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
我们可以不考虑输出结果的顺序。
进阶:

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小很多，哪种方法更优？
如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
*/
func intersect(nums1 []int, nums2 []int) []int {
	memo := make(map[int]int)

	for _, v := range nums1 {
		if _, ok := memo[v]; ok {
			memo[v]++
		} else {
			memo[v] = 1
		}
	}

	res := []int{}
	for _, num := range nums2 {
		if cnt, ok := memo[num]; ok && cnt > 0 {
			res = append(res, num)
			memo[num]--
		}
	}
	return res
}

// 给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，
// 使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。

// 示例 1:

// 输入: nums = [1,2,3,1], k = 3
// 输出: true
// 示例 2:

// 输入: nums = [1,0,1,1], k = 1
// 输出: true
// 示例 3:

// 输入: nums = [1,2,3,1,2,3], k = 2
// 输出: false
func containsNearbyDuplicate(nums []int, k int) bool {
	memo := make(map[int]int)
	for i, v := range nums {
		if idx, ok := memo[v]; ok {
			if i-idx <= k {
				return true
			}
		}
		memo[v] = i
	}
	return false
}

/*
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序。
*/
type byteSlice []byte

func (s byteSlice) Len() int           { return len(s) }
func (s byteSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s byteSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func groupAnagrams(strs []string) [][]string {
	memo := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		convertStr := byteSlice(strs[i])
		sort.Sort(convertStr)
		sortedStr := string(convertStr)
		if _, ok := memo[sortedStr]; !ok {
			memo[sortedStr] = make([]string, 0)
		}
		memo[sortedStr] = append(memo[sortedStr], strs[i])
	}

	ret := [][]string{}
	for _, v := range memo {
		ret = append(ret, v)
	}
	return ret
}

/*
给定一个字符串，对该字符串可以进行 “移位” 的操作，也就是将字符串中每个字母都变为其在字母表中后续的字母，
比如："abc" -> "bcd"。这样，我们可以持续进行 “移位” 操作，从而生成如下移位序列：

"abc" -> "bcd" -> ... -> "xyz"
给定一个包含仅小写字母字符串的列表，将该列表中所有满足 “移位” 操作规律的组合进行分组并返回。

示例：

输入: ["abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"]
输出:
[
  ["abc","bcd","xyz"],
  ["az","ba"],
  ["acef"],
  ["a","z"]
]
*/
func calShiftCharString(str string, cnt int) string {
	res := []byte(str)
	for i := 0; i < len(res); i++ {
		res[i] = ((res[i]-'a')+byte(cnt))%26 + 'a'
	}
	return string(res)
}
func groupStrings(strings []string) [][]string {
	memo := make(map[string][]string)
	for _, str := range strings {
		for i := 0; i < 26; i++ {
			shiftStr := calShiftCharString(str, i)
			if _, ok := memo[shiftStr]; ok {
				memo[shiftStr] = append(memo[shiftStr], str)
				goto NEXT
			}
		}

		memo[str] = []string{str}
	NEXT:
	}

	ret := [][]string{}
	for _, v := range memo {
		ret = append(ret, v)
	}
	return ret
}


/*
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。

上图是一个部分填充的有效的数独。

数独部分空格内已填入了数字，空白格用 '.' 表示。

示例 1:

输入:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: true
示例 2:

输入:
[
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: false
解释: 除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。
     但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。
说明:

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
给定数独序列只包含数字 1-9 和字符 '.' 。
给定数独永远是 9x9 形式的。
*/
func isValidSudoku(board [][]byte) bool {

}