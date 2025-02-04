// A mouse is trying to get from its starting position S to the target T,
// while moving only on land cells and staying as far away as possible
// from the cat C. Formally assume again a square grid of size (N*N),
// allowing only horizontal and vertical moves, impassable cells with water,
// and cells S, T and C. We want to find a path from S to T for which the
// minimal distance to C along the path is maximal. Use the L1 (aka Manhatton)
// distance measure, i.e, the sum of the horizontal and vertical distances.

// Mouse Path finding Problem: Avoiding the Cat 🐱
// Given an N × N grid with:

// S (Mouse's Start Position)
// T (Target Position)
// C (Cat's Position)
// Water Cells (Impassable)
// Allowed Moves: Only horizontal and vertical (no diagonals)
// Objective: Find a path from S to T that stays as far away
// from C as possible, i.e., maximize the minimum distance from C
// along the path.

// Solution Approach
// This problem requires a modified BFS (Breadth-First Search) approach:

// Precompute the Distance from the Cat (C):
// Use multi-source BFS to compute the shortest L1 (Manhattan) distance from C to every land cell.
// This results in a "danger heatmap".
// Perform a BFS from S to T, prioritizing cells that maximize
// the minimum distance from C along the path.

// Complexity Analysis
// Cat Distance Preprocessing (computeCatDistance) → O(N²)
// A single BFS traverses all reachable cells.
// Path finding (findSafestPath) → O(N² log N)
// Uses a priority queue with at most N² elements.
// Each insertion takes O(log N²) = O(log N).

// Overall Complexity:
// 𝑂(𝑁^2 log⁡𝑁)
// O(N^2 logN)
// This is efficient for large grids.

package main

import (
	"container/heap"
	"fmt"
)

// Directions for moving in 4 possible directions (up, down, left, right)
var directions = [][]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

type Node struct {
	row, col, minDistToCat, distToTarget int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].minDistToCat > pq[j].minDistToCat }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Compute shortest distance from C to all land cells using BFS
func computeCatDistance(grid [][]int, cat []int) [][]int {
	N := len(grid)
	dist := make([][]int, N)
	for i := range dist {
		dist[i] = make([]int, N)
		for j := range dist[i] {
			dist[i][j] = -1 // -1 means unreachable
		}
	}

	queue := [][]int{cat}
	dist[cat[0]][cat[1]] = 0

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			newRow, newCol := cell[0]+dir[0], cell[1]+dir[1]
			if newRow >= 0 && newRow < N && newCol >= 0 && newCol < N &&
				grid[newRow][newCol] == 0 && dist[newRow][newCol] == -1 {

				dist[newRow][newCol] = dist[cell[0]][cell[1]] + 1
				queue = append(queue, []int{newRow, newCol})
			}
		}
	}
	return dist
}

// Modified BFS to find the safest path from S to T
func findSafestPath(grid [][]int, start, target, cat []int) bool {
	N := len(grid)
	catDist := computeCatDistance(grid, cat)

	// Priority queue (max-heap) for exploring paths that maximize min distance from C
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Node{start[0], start[1], catDist[start[0]][start[1]], 0})

	visited := make(map[[2]int]bool)
	visited[[2]int{start[0], start[1]}] = true

	for pq.Len() > 0 {
		node := heap.Pop(pq).(*Node)

		// If we reached T, return true
		if node.row == target[0] && node.col == target[1] {
			fmt.Println("Safest Path Found! Maximum min-distance from Cat:", node.minDistToCat)
			return true
		}

		// Explore in all 4 directions
		for _, dir := range directions {
			newRow, newCol := node.row+dir[0], node.col+dir[1]

			if newRow >= 0 && newRow < N && newCol >= 0 && newCol < N &&
				grid[newRow][newCol] == 0 && !visited[[2]int{newRow, newCol}] {

				newMinDistToCat := min(node.minDistToCat, catDist[newRow][newCol])
				heap.Push(pq, &Node{newRow, newCol, newMinDistToCat, node.distToTarget + 1})
				visited[[2]int{newRow, newCol}] = true
			}
		}
	}

	fmt.Println("No Safe Path Found!")
	return false
}

func main() {
	// Example Grid (0 = land, 1 = water)
	grid := [][]int{
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 1, 0},
	}

	start := []int{0, 0}  // Mouse's start position (S)
	target := []int{4, 4} // Target (T)
	cat := []int{2, 2}    // Cat's position (C)

	findSafestPath(grid, start, target, cat)
}
