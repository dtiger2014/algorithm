package hash

type MyHashMap struct {
	Keys [][]int
	Vals [][]int
}

const (
	hashSize = 1000
)

/** Initialize your data structure here. */
func NewMyHastMap() MyHashMap {
	return MyHashMap{
		Keys: make([][]int, hashSize),
		Vals: make([][]int, hashSize),
	}
}

func (this *MyHashMap) hashFunc(key int) int {
	return key % hashSize
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	k := this.hashFunc(key)

	for i, v := range this.Keys[k] {
		if key == v {
			this.Vals[k][i] = value
			return
		}
	}

	this.Keys[k] = append(this.Keys[k], key)
	this.Vals[k] = append(this.Vals[k], value)
	return
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	k := this.hashFunc(key)
	for i, v := range this.Keys[k] {
		if key == v {
			return this.Vals[k][i]
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	k := this.hashFunc(key)

	for i, v := range this.Keys[k] {
		if key == v {
			if i == 0 {
				this.Keys[k] = this.Keys[k][1:]
				this.Vals[k] = this.Vals[k][1:]
			} else if i == len(this.Keys[k])-1 {
				this.Keys[k] = this.Keys[k][:len(this.Keys[k])-1]
				this.Vals[k] = this.Vals[k][:len(this.Vals[k])-1]
			} else {
				this.Keys[k] = append(this.Keys[k][:i], this.Keys[k][i+1:]...)
				this.Vals[k] = append(this.Vals[k][:i], this.Vals[k][i+1:]...)
			}
			return
		}
	}
	return
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
