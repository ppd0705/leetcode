package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	post := make(map[int][]int)
	preNum := make(map[int]int)

	for _, req := range prerequisites {
		a, b := req[0], req[1]
		post[b] = append(post[b], a)
		preNum[a]++
	}

	que := make([]int, 0)
	res := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if preNum[i] == 0 {
			que = append(que, i)
		}
	}
	for len(que) > 0 {
		l := len(que)
		for i := 0; i < l; i++ {
			res = append(res, que[i])
			for _, j := range post[que[i]] {
				preNum[j]--
				if preNum[j] == 0 {
					que = append(que, j)
				}
			}
		}
		que = que[l:]
	}
	if len(res) != numCourses {
		return []int{}
	}
	return res
}
