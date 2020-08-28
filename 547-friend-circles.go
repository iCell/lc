package main

func findCircleNumBFS(M [][]int) int {
	if len(M) == 0 {
		return 0
	}

	visited := make(map[int]bool, len(M))
	queue := make([]int, 0, len(M))

	var count int
	for i := range M {
		if visited[i] == true {
			continue
		}

		queue = append(queue, i)
		for len(queue) != 0 {
			pop, temp := queue[0], queue[1:]
			queue = temp

			for j, v := range M[pop] {
				if visited[j] == false && v == 1 {
					queue = append(queue, j)
					visited[j] = true
				}
			}
		}

		count++
	}

	return count
}

func findCircleNum(M [][]int) int {
	uf := NewUnionFind(len(M))

	for i := 0; i < len(M); i++ {
		for j := i; j < len(M[i]); j++ {
			if M[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}

	return uf.Size
}

type UnionFind struct {
	Size    int
	Parents []int
}

func NewUnionFind(size int) *UnionFind {
	parents := make([]int, size, size)
	for i := range parents {
		parents[i] = i
	}
	return &UnionFind{
		Size:    size,
		Parents: parents,
	}
}

func (uf *UnionFind) Union(l, r int) {
	l = uf.Find(l)
	r = uf.Find(r)
	if l == r {
		return
	}
	uf.Parents[l] = r
	uf.Size--
}

func (uf *UnionFind) Find(p int) int {
	for p != uf.Parents[p] {
		uf.Parents[p] = uf.Parents[uf.Parents[p]]
		p = uf.Parents[p]
	}
	return p
}
