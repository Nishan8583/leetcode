/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

    For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.

Return true if you can finish all courses. Otherwise, return false.

Solution:
Basically we need to check if there is acyclic node
Ex; [[1,0]] 1 requires 0 to be completed, thats all 0->1
Ex: [[1,0],[0,1]] 1 reuqires 0 to be completed but 0 also requires 1 to be completed, so it is a acyclic node. 0->1->0

we create a map of prerequisites courses to course {0:[1]}, this is our adjancency adj_list
After that We loop through our adj_list, put the current visited node in a temp path, and check if a node has been visited twice
for example: {0:[1]}, 0 will have only been visited twice, but for {0:[1],1:[0]}, 0 will have been twice visited and in this case, we return false


*/

package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	// map of prerequisites course to the actual course
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

	// fmt.Println(adj_list)
	// path will hold node visited to 1, to indicate that this path was alreay taken
	// basically acyclic node detection
	path := map[int]int{}

	// visited
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

// example for [1,0]
// dfs(0,{0:1},{},{})
func dfs(course int, adj_list map[int][]int, path map[int]int, visited map[int]int) bool {
	defer delete(path, course)

	// example c will be [1]
	c, ok := adj_list[course]

	// it has no value, i.e. this is not a prerequisites to any course
	if !ok {
		return true
	}

	// if it was already in current path, means its pointing towards itself, return false
	// fmt.Println("Checking for", course, path)
	_, ac := path[course]
	// fmt.Println("Checking for", course, path, ac)
	if ac {
		return false
	}
	// fmt.Println("adding", course)
	path[course] = 1
	visited[course] = 1

	// LOOPING through the course that require the "course" as a prerequisites
	// example [1]
	for _, neigbhour := range c {

		// if the node was already visited in the current path acyclic node
		_, ac := path[neigbhour]
		// fmt.Println("Checking for", course, path, ac)
		if ac {
			return false
		}

		if _, ok := visited[neigbhour]; ok {
			// fmt.Println("Not Visting", neigbhour)
			continue
		}
		// fmt.Println("Visiting", neigbhour)
		if v := dfs(neigbhour, adj_list, path, visited); !v {
			return v
		}
	}
	return true
}
