type MyHashSet struct {
	keySpace int
	buckets []*Bucket
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
	space := 2049
	buckets := make([]*Bucket, 0, space)
	for i := 0; i < space; i++ {
		buckets = append(buckets, NewBucket())
	}
	return MyHashSet{
		keySpace: space,
		buckets: buckets,
	}
}


func (this *MyHashSet) Add(key int)  {
	hashKey := key % this.keySpace
	this.buckets[hashKey].Update(key)
}


func (this *MyHashSet) Remove(key int)  {
	hashKey := key % this.keySpace
	this.buckets[hashKey].Delete(key)
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	hashKey := key % this.keySpace
	return this.buckets[hashKey].Get(key) != -1
}


type Bucket struct {
	nodes []int
}

func NewBucket() *Bucket {
	return &Bucket{nodes: make([]int, 0)}
}

func (b *Bucket) Get(key int) int {
	for _, node := range b.nodes {
		if node == key {
			return node
		}
	}
	return -1
}

func (b *Bucket) Update(key int) {
	for i, node := range b.nodes {
		if node == key {
			b.nodes[i] = key
			return
		}
	}
	b.nodes = append(b.nodes, key)
}

func (b *Bucket) Delete(key int) {
	idx := -1
	for i, node := range b.nodes {
		if node == key {
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
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */