package main

func minMutation(start string, end string, bank []string) int {
	if start == end {
		return 0
	}

	charset := []rune{'A', 'C', 'G', 'T'}

	genetics := make(map[string]bool)
	for _, g := range bank {
		genetics[g] = true
	}

	level := 0
	queue := &Queue{}
	queue.InQueue(start)

	for !queue.IsEmpty() {
		size := queue.Size()
		for size > 0 {
			cur := queue.DeQueue()
			if cur == end {
				return level
			}

			runes := []rune(cur)
			for i, _ := range runes {
				old := runes[i]
				for _, c := range charset {
					runes[i] = c
					next := string(runes)

					_, ok := genetics[next]
					if !ok {
						continue
					}
					queue.InQueue(next)
					delete(genetics, next)
				}
				runes[i] = old
			}
			size--
		}
		level++
	}

	return -1
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

func minMutation2(start string, end string, bank []string) int {
	if start == end {
		return 0
	}

	genetics := make(map[string]bool)
	for _, b := range bank {
		genetics[b] = true
	}

	_, valid := genetics[end]
	if !valid {
		return -1
	}

	charset := []rune{'A', 'C', 'G', 'T'}

	level := 0
	starts, ends := make(map[string]bool), make(map[string]bool)
	starts[start] = true
	ends[end] = true

	for len(starts) > 0 && len(ends) > 0 {
		if len(starts) > len(ends) {
			starts, ends = ends, starts
		}

		temp := make(map[string]bool)

		for gen := range starts {
			runes := []rune(gen)
			for i, old := range runes {
				for _, c := range charset {
					if c == old {
						continue
					}
					runes[i] = c

					next := string(runes)

					_, exist := ends[next]
					if exist {
						return level + 1
					}

					_, ok := genetics[next]
					if !ok {
						continue
					}

					delete(genetics, next)
					temp[next] = true
				}
				runes[i] = old
			}
		}

		starts = temp
		level++
	}

	return -1
}
