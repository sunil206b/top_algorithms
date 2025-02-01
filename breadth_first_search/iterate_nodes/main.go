package main

import "fmt"

func main() {
	graph := map[string][]string{
		"A": {"B", "G"},
		"B": {"C", "D", "E"},
		"C": {},
		"D": {},
		"E": {"F"},
		"F": {},
		"G": {"H"},
		"H": {"I"},
		"I": {},
	}

	// fmt.Println(">>>> DFL recursive traversal <<<<<")
	// printNodesRecursively(graph, "A")

	fmt.Println()
	fmt.Println(">>>> DFL iterative traversal <<<<<")
	printNodesIteratively(graph, "A")
}

func printNodesIteratively(graph map[string][]string, startNode string) {
	visited := make(map[string]bool)
	queue := []string{startNode}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		visited[node] = true
		fmt.Print(node, " ")

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
}
