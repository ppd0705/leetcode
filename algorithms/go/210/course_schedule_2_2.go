package main

func findOrder2(numCourses int, prerequisites [][]int) []int {
	post := make(map[int][]int)
	res := make([]int, 0)
	visited := make([]int, numCourses)

	for _, req := range prerequisites {
		a, b := req[0], req[1]
		post[b] = append(post[b], a)
	}
	var dfs func(u int)
	valid := true
	dfs = func(u int) {
		visited[u] = 1
		for _, v := range post[u] {
			if visited[v] == 1 {
				valid = false
				return
			} else if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			}
		}
		visited[u] = 2
		res = append(res, u)
	}
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if !valid {
		return []int{}
	}
	for i := 0; i < numCourses/2; i++ {
		res[i], res[numCourses-i-1] = res[numCourses-i-1], res[i]
	}
	return res
}
