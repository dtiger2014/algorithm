package hash

type MyHashSet struct {
	data [][]int
}

const (
	hashSize = 1000
)

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	ret := MyHashSet{
		data: make([][]int, 1000),
	}

	for i := range ret.data {
		ret.data[i] = make([]int, 0)
	}

	return ret
}

func (this *MyHashSet) hashFunc(value int) int {
	key1 := value % hashSize

	return key1
}

func (this *MyHashSet) Add(key int) {
	k := this.hashFunc(key)

	if len(this.data[k]) == 0 {
		this.data[k] = append(this.data[k], key)
	} else {
		find := false
		for i := 0; i < len(this.data[k]); i++ {
			if this.data[k][i] == key {
				find = true
				break
			}
		}
		if find == false {
			this.data[k] = append(this.data[k], key)
		}
	}
}

func (this *MyHashSet) Remove(key int) {
	k := this.hashFunc(key)

	for i, v := range this.data[k] {
		if v == key {
			if i == 0 {
				this.data[k] = this.data[k][1:]
			} else if i == len(this.data[k])-1 {
				this.data[k] = this.data[k][:len(this.data[k])-1]
			} else {
				this.data[k] = append(this.data[k][:i], this.data[k][i+1:]...)
			}
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	k := this.hashFunc(key)
	for _, v := range this.data[k] {
		if v == key {
			return true
		}
	}

	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
