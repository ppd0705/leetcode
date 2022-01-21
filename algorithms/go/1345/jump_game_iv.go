package main

func minJumps(arr []int) int {
	visited := map[int]bool{0: true}
	queue := []int{0}
	neighbors := make(map[int][]int)
	for i, v := range arr {
		neighbors[v] = append(neighbors[v], i)
	}
	step := 0
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			idx := queue[i]
			if idx == len(arr)-1 {
				return step
			}
			for _, neighbor := range neighbors[arr[idx]] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
					visited[neighbor] = true
				}
			}
			delete(neighbors, arr[idx])
			if idx+1 < len(arr) && !visited[idx+1] {
				queue = append(queue, idx+1)
				visited[idx+1] = true
			}
			if idx > 0 && !visited[idx-1] {
				queue = append(queue, idx-1)
				visited[idx-1] = true
			}
		}
		step++
		queue = queue[l:]
	}
	return -1
}
