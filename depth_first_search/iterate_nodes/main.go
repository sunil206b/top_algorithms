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

	fmt.Println(">>>> DFL recursive traversal <<<<<")
	printNodesRecursively(graph, "A")

	fmt.Println()
	fmt.Println(">>>> DFL iterative traversal <<<<<")
	printNodesIteratively(graph, "A")
}

func printNodesRecursively(graph map[string][]string, startNode string) {
	visited := make(map[string]bool)

	var dfs func(node string)
	dfs = func(node string) {
		visited[node] = true
		fmt.Print(node, " ")

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	dfs(startNode)
}

func printNodesIteratively(graph map[string][]string, startNode string) {
	visited := make(map[string]bool)
	stack := []string{startNode}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[node] {
			visited[node] = true
			fmt.Print(node, " ")

			for _, neighbor := range graph[node] {
				if !visited[neighbor] {
					stack = append(stack, neighbor)
				}
			}
		}
	}
}
