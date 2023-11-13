package main

/*
https://leetcode.com/problems/course-schedule/submissions/
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false.

Its a topological sort
convert prerequisites into an adjencancy list, then follow the topological sort algorithm
 topological sort algorithm: Do a dfs, In this case, it is a bit easier, because we just to need to check if there is acyclical
 We do a DFS, a path map, once that path has been completed we pop it out, but if we find it in the path again, we return false
 visited in outer most loop and in inner most to prevent going through same path again

*/
import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj_list := make(map[int][]int)

	if len(prerequisites) == 0 {
		return true
	}
	for _, course := range prerequisites {
		if course[1] == course[0] {
			return false
		}
		if v, ok := adj_list[course[1]]; ok {
			adj_list[course[1]] = append(v, course[0])
			continue
		}
		adj_list[course[1]] = []int{course[0]}
	}

	//fmt.Println(adj_list)
	// path will hold node visited to 1, to indicate that this path was alreay taken
	// basically acyclic node detection
	path := map[int]int{}

	visited := map[int]int{}
	for _, v := range prerequisites {
		if _, ok := visited[v[1]]; ok {
			continue
		}
		if !dfs(v[1], adj_list, path, visited) {
			return false
		}
	}
	return true
}

func dfs(course int, adj_list map[int][]int, path map[int]int, visited map[int]int) bool {
	defer delete(path, course)

	c, ok := adj_list[course]

	// it has no value
	if !ok {
		return true
	}

	// if it was already in path, means its pointing towards itself, return false
	//fmt.Println("Checking for", course, path)
	_, ac := path[course]
	//fmt.Println("Checking for", course, path, ac)
	if ac {
		return false
	}
	//fmt.Println("adding", course)
	path[course] = 1
	visited[course] = 1
	for _, neigbhour := range c {
		_, ac := path[neigbhour]
		//fmt.Println("Checking for", course, path, ac)
		if ac {
			return false
		}
		if _, ok := visited[neigbhour]; ok {
			//fmt.Println("Not Visting", neigbhour)
			continue
		}
		//fmt.Println("Visiting", neigbhour)
		if v := dfs(neigbhour, adj_list, path, visited); !v {
			return v
		}
	}
	return true
}

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(canFinish(3, [][]int{{0, 2}, {1, 2}, {2, 0}}))
	fmt.Println(canFinish(3, [][]int{{1, 0}, {2, 1}, {3, 4}, {4, 3}}))
}
