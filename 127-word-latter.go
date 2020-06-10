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
