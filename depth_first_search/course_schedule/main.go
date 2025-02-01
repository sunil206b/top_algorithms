package main

// There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
// You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that
// you must take course bi first if you want to take course ai.

// For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
// Return true if you can finish all courses. Otherwise, return false.

// Example 1:
// Input: numCourses = 2, prerequisites = [[1,0]]
// Output: true
// Explanation: There are a total of 2 courses to take.
// To take course 1 you should have finished course 0. So it is possible.

// Example 2:
// Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
// Output: false
// Explanation: There are a total of 2 courses to take.
// To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.

import "fmt"

func main() {
	// Example Test Cases
	numCourses1 := 2
	prerequisites1 := [][]int{{1, 0}}                                          // 0 -> 1 (Can complete all)
	fmt.Println("Can finish courses?", canFinish(numCourses1, prerequisites1)) // true

	numCourses2 := 2
	prerequisites2 := [][]int{{1, 0}, {0, 1}}                                  // 0 -> 1 -> 0 (Cycle)
	fmt.Println("Can finish courses?", canFinish(numCourses2, prerequisites2)) // false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	// Step 1: Create adjacency list
	graph := make(map[int][]int)

	for _, pre := range prerequisites {
		course, prerequisite := pre[0], pre[1]
		graph[prerequisite] = append(graph[prerequisite], course)
	}

	// Step 2: States: 0 -> unvisited, 1 -> visiting, 2 -> visited
	visited := make(map[int]int)

	// DFS function to detect cycles
	var dfs func(course int) bool
	dfs = func(course int) bool {
		// Cycle detected
		if visited[course] == 1 {
			return true
		}

		// Already processed
		if visited[course] == 2 {
			return false
		}

		// Mark as visiting (1)
		visited[course] = 1

		// Visit all prerequisites (neighbors)
		for _, neighbor := range graph[course] {
			if dfs(neighbor) {
				return true
			}
		}

		// Mark as fully processed (2)
		visited[course] = 2
		return false
	}

	// Step 3: Check all courses for cycles
	for course := 0; course < numCourses; course++ {
		if visited[course] == 0 {
			if dfs(course) {
				return false // Cycle detected, cannot finish all courses
			}
		}
	}

	return true // No cycle found, can finish all courses
}
