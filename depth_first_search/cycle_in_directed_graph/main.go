package main

import "fmt"

func main() {
	// Example graph with a cycle: 0 → 1 → 2 → 0
	graphWithCycle := map[int][]int{
		0: {1},
		1: {2},
		2: {0},
	}

	// Example graph without a cycle: 0 → 1 → 2
	graphWithoutCycle := map[int][]int{
		0: {1},
		1: {2},
		2: {},
	}

	fmt.Println("Graph with cycle:", hasCycle(graphWithCycle))
	fmt.Println("Graph without cycle:", hasCycle(graphWithoutCycle))
}

func hasCycle(graph map[int][]int) bool {
	// States: 0 -> unvisited, 1 -> visiting, 2 -> visited
	visited := make(map[int]int)

	// Helper function for DFS
	var dfs func(node int) bool
	dfs = func(node int) bool {
		// If the node is being visited,we have a cycle
		if visited[node] == 1 {
			return true
		}

		// If the node is fully visited, no need to process again
		if visited[node] == 2 {
			return false
		}

		// Mark the node is being visited
		visited[node] = 1

		// Recur for all neighbors
		for _, neighbor := range graph[node] {
			if dfs(neighbor) {
				return true
			}
		}

		// Mark the node is fully visited
		visited[node] = 2
		return false
	}

	// Run DFS for all nodes in the graph
	for node := range graph {
		if visited[node] == 0 {
			if dfs(node) {
				return true
			}
		}
	}

	return false
}
