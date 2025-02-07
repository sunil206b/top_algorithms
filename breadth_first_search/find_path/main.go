package main

import "fmt"

func main() {
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
	queue := []string{src}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		visited[node] = true

		if node == dst {
			return true
		}

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
	return false
}
