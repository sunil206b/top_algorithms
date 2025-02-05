package main

import (
	"container/list"
	"fmt"
)

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println("DFS -> Islands count: ", dfsNumberOfIslands(grid))
	fmt.Println("BFS -> Islands count: ", bfsNumberOfIslands(grid))

	grid1 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println("DFS -> Islands count: ", dfsNumberOfIslands(grid1))
	fmt.Println("BFS -> Islands count: ", bfsNumberOfIslands(grid1))
}

func dfsNumberOfIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	newGrid := deepCopy2D(grid)

	rows, cols := len(newGrid), len(newGrid[0])
	count := 0
	var dfs func(r, c int)
	dfs = func(r, c int) {
		// Boundary check and water check
		if r < 0 || r >= rows || c < 0 || c >= cols || newGrid[r][c] == '0' {
			return
		}

		// Mark current land as visited (turn '1' to '0')
		newGrid[r][c] = '0'

		// Visit all 4 directions
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if newGrid[r][c] == '1' { // Found an unvisited island
				count++
				dfs(r, c) // Mark all connected land
			}
		}
	}

	return count
}

func bfsNumberOfIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	newGrid := deepCopy2D(grid)

	rows, cols := len(newGrid), len(newGrid[0])
	count := 0

	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if newGrid[r][c] == '1' {
				count++
				queue := list.New()
				queue.PushBack([]int{r, c})
				newGrid[r][c] = '0' // Mark as visited

				for queue.Len() > 0 {
					node := queue.Remove(queue.Front()).([]int)
					row, col := node[0], node[1]

					for _, dir := range directions {
						newRow, newCol := row+dir[0], col+dir[1]

						if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && newGrid[newRow][newCol] == '1' {
							queue.PushBack([]int{newRow, newCol})
							newGrid[newRow][newCol] = '0' // Mark as visited
						}
					}
				}
			}
		}
	}

	return count
}

func deepCopy2D(original [][]byte) [][]byte {
	newList := make([][]byte, len(original)) // Create outer slice

	for i := range original {
		newList[i] = make([]byte, len(original[i])) // Create inner slice
		copy(newList[i], original[i])               // Copy elements
	}

	return newList
}
