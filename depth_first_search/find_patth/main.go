package main

import "fmt"

func main() {
	graph := map[string][]string{
		"f": {"g", "i"},
		"g": {"h"},
		"h": {},
		"i": {"g", "k"},
		"j": {"i"},
		"k": {},
	}

	fmt.Println(hasPath(graph, "f", "k"))
	fmt.Println(hasPath(graph, "f", "j"))
	fmt.Println(hasPath(graph, "i", "h"))

	// Directed Acyclic Graph (DAG)
	// Graph Structure (DAG):
	//     A
	//    / \
	//   B   C
	//  / \   \
	// D   E   F
	// |  / \ / \
	// G  G  H  I
	//  \ |  |  |
	//   \|  |  |
	//    J  J  J
	graph1 := map[string][]string{
		"A": {"B", "C"},
		"B": {"D", "E"},
		"C": {"F"},
		"D": {"G"},
		"E": {"G", "H"},
		"F": {"H", "I"},
		"G": {"J"},
		"H": {"J"},
		"I": {"J"},
		"J": {},
	}
	fmt.Println(hasPath(graph1, "A", "J"))
	fmt.Println(hasPath(graph1, "B", "F"))
	fmt.Println(hasPath(graph1, "B", "H"))
}

func hasPath(graph map[string][]string, src, dst string) bool {
	visited := make(map[string]bool)

	var dfs func(node string) bool
	dfs = func(node string) bool {
		visited[src] = true
		if node == dst {
			return true
		}

		for _, neighbor := range graph[node] {
			if !visited[neighbor] && dfs(neighbor) {
				return true
			}
		}

		return false
	}

	return dfs(src)
}
