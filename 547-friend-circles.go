package main

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
