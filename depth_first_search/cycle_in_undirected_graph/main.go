package main

import "fmt"

func main() {
	// Example graph with a cycle: 0 - 1 - 2 - 0
	graphWithCycle := map[int][]int{
		0: {1, 2},
		1: {0, 2},
		2: {0, 1},
	}

	// Example graph without a cycle: 0 - 1 - 2
	graphWithoutCycle := map[int][]int{
		0: {1},
		1: {0, 2},
		2: {1},
	}

	fmt.Println("Graph with cycle:", hasCycle(graphWithCycle))       // Output: true
	fmt.Println("Graph without cycle:", hasCycle(graphWithoutCycle)) // Output: false
}

func hasCycle(graph map[int][]int) bool {
	visited := make(map[int]bool)

	var dfs func(node, parent int) bool
	dfs = func(node, parent int) bool {
		visited[node] = true

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				if dfs(neighbor, node) {
					return true
				}
			} else if neighbor != parent {
				return true
			}
		}

		return false
	}

	for node := range graph {
		if !visited[node] {
			if dfs(node, -1) {
				return true
			}
		}
	}

	return false
}
