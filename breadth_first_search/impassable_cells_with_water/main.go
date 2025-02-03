// You are given an N × N grid where:
// S (Start Cell) and T (Target Cell) are given.
// Some cells contain water (impassable).
// You can only move horizontally or vertically (no diagonal moves).
// The goal is to determine if a path exists from S to T.

// Solution Approach:
// Since this is a path finding problem in an unweighted grid, we can use Breadth-First Search (BFS) because:

// BFS guarantees the shortest path in an unweighted graph.
// It efficiently checks if T is reachable from S without getting stuck in water cells.

package main

import "fmt"

// Directions for moving in 4 possible directions (up, down, left, right)
var directions = [][]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func main() {
	// Example Grid (0 = passable, 1 = water)
	grid := [][]int{
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 1, 0},
	}

	start := []int{0, 0}  // S at (0,0)
	target := []int{4, 4} // T at (4,4)

	if isPathExist(grid, start, target) {
		fmt.Println("Path Exists!")
	} else {
		fmt.Println("No Path Found!")
	}
}

func isPathExist(grid [][]int, start []int, target []int) bool {
	// Length of the grid
	N := len(grid)

	// If start or target is in water, return false
	if grid[start[0]][start[1]] == 1 || grid[target[0]][target[1]] == 1 {
		return false
	}

	// To track the visited cells
	visited := make(map[[2]int]bool)
	// Queue for BFS
	queue := [][]int{start}
	// Mark start cell is visited
	visited[[2]int{start[0], start[1]}] = true

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		if cell[0] == target[0] && cell[1] == target[1] {
			return true
		}

		for _, dir := range directions {
			newRow, newCol := cell[0]+dir[0], cell[1]+dir[1]

			// Check if new position is valid
			if newRow > 0 && newRow < N && newCol >= 0 && newCol < N &&
				grid[newRow][newCol] == 0 && !visited[[2]int{newRow, newCol}] {

				visited[[2]int{newRow, newCol}] = true
				queue = append(queue, []int{newRow, newCol})
			}
		}
	}

	// If BFS completes without finding T, return false
	return false
}
