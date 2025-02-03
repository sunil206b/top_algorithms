// When dealing with unweighted graphs, BFS (Breadth-First Search) is the best approach
// for finding the shortest path from a single source node to all other nodes.

// Why BFS Works for Shortest Path in Unweighted Graphs?
// > BFS explores all nodes level by level.
// > The first time a node is reached, it is at its shortest distance from the source.
// > It processes each node exactly once, making it efficient for shortest path calculations.

package main

import (
	"fmt"
	"math"
)

func main() {
	// Graph representation (Adjacency List)
	// 	(10)----(1)
	//    |      |
	//    |      |      (12)----(2)
	//    |      |       |       |
	//   (9)----(8)------|      (3)----(4)
	//    |                      |
	//    |                      |
	//   (0)-----|               |
	//    |      |               |
	//    |     (7)--------------|
	//    |      | \
	//   (11)----|  \
	//              (6)-----(5)
	graph := map[int][]int{
		0:  {7, 9, 11},
		1:  {10, 8},
		2:  {3, 12},
		3:  {2, 4},
		4:  {3},
		5:  {6},
		6:  {5, 7},
		7:  {0, 3, 6, 11},
		8:  {1, 9, 12},
		9:  {0, 8, 10},
		10: {1, 9},
		11: {0, 7},
		12: {2, 8},
	}

	startNode := 10
	targetNode := 5
	shortestPaths := BFSShortestDistances(graph, startNode)

	fmt.Println("Shortest distances from node", startNode)
	for node, distance := range shortestPaths {
		fmt.Printf("Node %d: %d\n", node, distance)
	}

	shortestPath := BFSShortestPath(graph, startNode, targetNode)

	// Output shortest path from start to target
	if shortestPath == nil {
		fmt.Printf("No path found from %d to %d\n", startNode, targetNode)
	} else {
		fmt.Printf("Shortest path from %d to %d: %v\n", startNode, targetNode, shortestPath)
	}
}

func BFSShortestDistances(graph map[int][]int, startNode int) map[int]int {
	dist := make(map[int]int)
	for node := range graph {
		dist[node] = math.MaxInt32
	}
	dist[startNode] = 0

	queue := []int{startNode}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbor := range graph[current] {
			if dist[neighbor] == math.MaxInt32 {
				dist[neighbor] = dist[current] + 1
				queue = append(queue, neighbor)
			}
		}
	}

	return dist
}

func BFSShortestPath(graph map[int][]int, start, target int) []int {
	dist := make(map[int]int)
	parent := make(map[int]int)

	for node := range graph {
		dist[node] = math.MaxInt32
	}
	dist[start] = 0

	queue := []int{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbor := range graph[current] {
			if dist[neighbor] == math.MaxInt32 {
				dist[neighbor] = dist[current] + 1
				queue = append(queue, neighbor)
				parent[neighbor] = current
			}
		}
	}

	if dist[target] == math.MaxInt32 {
		return nil // Return nil if no path found
	}

	path := []int{}
	// Backtrack the shortest path from the target
	for at := target; at != start; at = parent[at] {
		path = append([]int{at}, path...) // Append node at the beginning
	}
	path = append([]int{start}, path...) // Add start node

	// If S and E are connected return the path.
	if path[0] == start {
		return path
	}

	return nil
}
