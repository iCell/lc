type MyHashMap struct {
	keySpace int
	buckets []*Bucket
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	space := 2049
	buckets := make([]*Bucket, 0, space)
	for i := 0; i < space; i++ {
		buckets = append(buckets, NewBucket())
	}
	return MyHashMap{
		keySpace: space,
		buckets: buckets,
	}
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
	hashKey := key % this.keySpace
	this.buckets[hashKey].Update(key, value)
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	hashKey := key % this.keySpace
	return this.buckets[hashKey].Get(key)
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
	hashKey := key % this.keySpace
	this.buckets[hashKey].Delete(key)
}

type Node struct {
	Key int
	Val int
}

func NewNode(key, val int) *Node {
	return &Node {
		Key: key,
		Val: val,
	}
}

type Bucket struct {
	nodes []*Node
}

func NewBucket() *Bucket {
	return &Bucket{nodes: make([]*Node, 0)}
}

func (b *Bucket) Get(key int) int {
	for _, node := range b.nodes {
		if node.Key == key {
			return node.Val
		}
	}
	return -1
}

func (b *Bucket) Update(key, val int) {
	for i, node := range b.nodes {
		if node.Key == key {
			b.nodes[i] = NewNode(key, val)
			return
		}
	}
	b.nodes = append(b.nodes, NewNode(key, val))
}

func (b *Bucket) Delete(key int) {
	idx := -1
	for i, node := range b.nodes {
		if node.Key == key {
			idx = i
			break
		}
	}    
	if idx == -1 {
		return
	}
	b.nodes[idx], b.nodes[len(b.nodes)-1] = b.nodes[len(b.nodes)-1], b.nodes[idx]
	b.nodes = b.nodes[:len(b.nodes)-1]
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */