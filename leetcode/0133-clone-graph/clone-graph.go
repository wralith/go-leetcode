package leetcode

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	return dfs(node, visited)

}

func dfs(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if val, ok := visited[node]; ok {
		return val
	}

	n := &Node{Val: node.Val, Neighbors: []*Node{}}
	visited[node] = n

	for _, n := range node.Neighbors {
		visited[node].Neighbors = append(visited[node].Neighbors, dfs(n, visited))
	}

	return visited[node]
}
