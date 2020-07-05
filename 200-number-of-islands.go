package main

var directions = [][]int{[]int{0, -1}, []int{-1, 0}}

func numIslands3(grid [][]byte) int {
	m, n := len(grid), 0
	if m == 0 {
		return 0
	}
	n = len(grid[0])
	if n == 0 {
		return 0
	}

	uf := NewUnionFind(m * n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == byte('1') {
				for _, direction := range directions {
					x, y := i+direction[0], j+direction[1]
					if x >= 0 && x < m && y >= 0 && y < n {
						if grid[x][y] == byte('1') {
							uf.Union(i*n+j, x*n+y)
						}
					}
				}
			}
		}
	}

	result := make(map[int]bool, uf.Size)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == byte('1') {
				result[uf.Find(i*n+j)] = true
			}
		}
	}

	return len(result)
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

func (uf *UnionFind) Find(p int) int {
	for p != uf.Parents[p] {
		uf.Parents[p] = uf.Parents[uf.Parents[p]]
		p = uf.Parents[p]
	}
	return p
}

func (uf *UnionFind) Union(p, q int) {
	l, r := uf.Find(p), uf.Find(q)
	if l == r {
		return
	}
	uf.Parents[l] = r
	uf.Size--
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	num := 0
	for row, rows := range grid {
		for column, value := range rows {
			if value == '0' {
				continue
			}
			dfs(row, column, grid)
			num++
		}
	}

	return num
}

func dfs(row int, column int, grid [][]byte) {
	if row < 0 || column < 0 || row >= len(grid) || column >= len(grid[0]) || grid[row][column] == '0' {
		return
	}
	grid[row][column] = '0'
	dfs(row+1, column, grid)
	dfs(row, column+1, grid)
	dfs(row-1, column, grid)
	dfs(row, column-1, grid)
}

func numIslands2(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	numr, numc := len(grid), len(grid[0])
	nums := 0
	queue := Queue{}

	for r, rows := range grid {
		for c, v := range rows {
			if v == byte('0') {
				continue
			}
			nums++
			grid[r][c] = byte('0')
			queue.InQueue([]int{r, c})
			for !queue.IsEmpty() {
				size := queue.Size()
				for size > 0 {
					p := queue.DeQueue()
					x, y := p[0], p[1]
					if x+1 < numr && grid[x+1][y] == byte('1') {
						grid[x+1][y] = byte('0')
						queue.InQueue([]int{x + 1, y})
					}
					if y+1 < numc && grid[x][y+1] == byte('1') {
						grid[x][y+1] = byte('0')
						queue.InQueue([]int{x, y + 1})
					}
					if x-1 >= 0 && grid[x-1][y] == byte('1') {
						grid[x-1][y] = byte('0')
						queue.InQueue([]int{x - 1, y})
					}
					if y-1 >= 0 && grid[x][y-1] == byte('1') {
						grid[x][y-1] = byte('0')
						queue.InQueue([]int{x, y - 1})
					}
					size--
				}
			}
		}
	}

	return nums
}

type Queue struct {
	values [][]int
}

func (q *Queue) InQueue(n []int) {
	q.values = append(q.values, n)
}

func (q *Queue) DeQueue() []int {
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
