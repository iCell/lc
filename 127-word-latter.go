package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 0
	}

	words := make(map[string]bool)
	for _, g := range wordList {
		words[g] = true
	}

	level := 0
	queue := &Queue{}
	queue.InQueue(beginWord)

	for !queue.IsEmpty() {
		size := queue.Size()
		for size > 0 {
			cur := queue.DeQueue()
			if cur == endWord {
				return level + 1
			}

			runes := []rune(cur)
			for i, _ := range runes {
				old := runes[i]
				for c := 'a'; c <= 'z'; c++ {
					if c == old {
						continue
					}
					runes[i] = c
					next := string(runes)

					_, ok := words[next]
					if !ok {
						continue
					}
					queue.InQueue(next)
					delete(words, next)
				}
				runes[i] = old
			}
			size--
		}
		level++
	}

	return 0
}

func ladderLength2(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 0
	}

	words := NewHashSet()
	for _, w := range wordList {
		words.Add(w)
	}

	length := 1

	beginSet, endSet, visited := NewHashSet(), NewHashSet(), NewHashSet()
	beginSet.Add(beginWord)
	endSet.Add(endWord)

	for !beginSet.IsEmpty() && !endSet.IsEmpty() {
		if beginSet.Size() > endSet.Size() {
			beginSet, endSet = endSet, beginSet
		}
		temp := NewHashSet()
		for _, v := range beginSet.Values() {
			runes := []rune(v)
			for i, old := range runes {
				for c := 'a'; c <= 'z'; c++ {
					if c == old {
						continue
					}
					runes[i] = c

					next := string(runes)
					if endSet.Contain(next) {
						return length + 1
					}
					if visited.Contain(next) || !words.Contain(next) {
						continue
					}
					temp.Add(next)
					visited.Add(next)
					runes[i] = old
				}
			}
		}
		beginSet = temp
		length++
	}

	return 0
}

type HashSet struct {
	values map[string]bool
}

func NewHashSet() *HashSet {
	return &HashSet{
		values: make(map[string]bool),
	}
}

func (set *HashSet) Add(s string) {
	set.values[s] = true
}

func (set *HashSet) Contain(s string) bool {
	_, ok := set.values[s]
	return ok
}

func (set *HashSet) IsEmpty() bool {
	return len(set.values) == 0
}

func (set *HashSet) Size() int {
	return len(set.values)
}

func (set *HashSet) Values() []string {
	var r []string
	for k, _ := range set.values {
		r = append(r, k)
	}
	return r
}

type Queue struct {
	values []string
}

func (q *Queue) InQueue(n string) {
	q.values = append(q.values, n)
}

func (q *Queue) DeQueue() string {
	if q.IsEmpty() {
		return ""
	}
	first, values := q.values[0], q.values[1:]
	q.values = values
	return first
}

func (q *Queue) IsEmpty() bool {
	return len(q.values) == 0
}

func (q *Queue) Size() int {
	return len(q.values)
}
