package main

type Node struct {
	Val      string
	Children []*Node
}

func (n *Node) foreach() {

}

type Graph struct {
	Root    *Node
	End     *Node
	HashMap map[string]*Node
}

func NewGraph(root, end string) *Graph {
	return &Graph{
		HashMap: make(map[string]*Node),
		Root:    &Node{Val: root},
		End:     &Node{Val: end},
	}
}

func (g *Graph) AddChild(parent string, child string) {
	p := g.HashMap[parent]
	p.Children = append(p.Children, &Node{Val: child})
}

func (g *Graph) Convert() [][]string {
	return [][]string{}
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	words := make(map[string]bool)
	for _, w := range wordList {
		words[w] = true
	}

	queue := &Queue{}
	queue.InQueue(beginWord)
	graph := NewGraph(beginWord, endWord)

	level := 0
	for !queue.IsEmpty() {
		size := queue.Size()

		for size > 0 {
			cur := queue.DeQueue()
			if cur == endWord {
				return graph.Convert()
			}

			runes := []rune(cur)
			for i, _ := range runes {
				old := runes[i]
				for c := 'a'; c <= 'z'; c++ {
					runes[i] = c
					next := string(runes)

					_, ok := words[next]
					if !ok {
						continue
					}
					queue.InQueue(next)
					delete(words, next)

					graph.AddChild(cur, next)
				}
				runes[i] = old
			}

			size--
		}
		level++
	}

	return [][]string{}
}
