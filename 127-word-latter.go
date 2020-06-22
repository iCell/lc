package main

func ladderLength3(beginWord string, endWord string, wordList []string) int {
	words := make(map[string]bool)
	for _, word := range wordList {
		words[word] = true
	}
	_, ok := words[endWord]
	if !ok {
		return 0
	}

	begins := make(map[string]bool)
	begins[beginWord] = true

	steps := 0
	for len(begins) > 0 {
		steps++

		// use this way to treat it like a queue
		temp := make(map[string]bool)
		for w := range begins {
			runes := []rune(w)
			for i, old := range runes {
				for c := 'a'; c <= 'z'; c++ {
					runes[i] = c
					next := string(runes)

					if next == endWord {
						return steps + 1
					}
					_, valid := words[next]
					if !valid {
						continue
					}
					delete(words, next)
					temp[next] = true
				}
				runes[i] = old
			}
		}
		begins = temp
	}

	return 0
}

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
	words := make(map[string]bool)
	for _, word := range wordList {
		words[word] = true
	}
	_, ok := words[endWord]
	if !ok {
		return 0
	}

	begins, ends := make(map[string]bool), make(map[string]bool)
	begins[beginWord] = true
	ends[endWord] = true

	steps := 0
	for len(begins) > 0 && len(ends) > 0 {
		steps++
		if len(begins) > len(ends) {
			begins, ends = ends, begins
		}

		temp := make(map[string]bool)
		for w := range begins {
			runes := []rune(w)
			for i, old := range runes {
				for c := 'a'; c <= 'z'; c++ {
					runes[i] = c
					next := string(runes)
					_, exist := ends[next]
					if exist {
						return steps + 1
					}
					_, valid := words[next]
					if !valid {
						continue
					}
					delete(words, next)
					temp[next] = true
				}
				runes[i] = old
			}
		}
		begins = temp
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
