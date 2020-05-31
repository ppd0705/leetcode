package main

type pointer struct {
	x int
	y int
}

func updateBoard(board [][]byte, click []int) [][]byte {
	row := len(board)
	if row == 0 {
		return board
	}
	column := len(board[0])
	x, y := click[0], click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	} else if board[x][y] != 'E' {
		return board
	}

	dx := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
	set := map[pointer]bool{pointer{x, y}:false}
	for len(set) > 0 {
		newSet := map[pointer]bool{}
		for p:= range set {
			mine := uint8(0)
			pointers := make([]pointer, 0)
			for i:=0;i< 8;i++ {
				x, y := p.x+dx[i], p.y+dy[i]
				if x >=0 && x < row && y >= 0 && y < column {
					if board[x][y] == 'M' {
						mine++
					} else if board[x][y] == 'E' {
						pointers = append(pointers, pointer{x, y})
					}
				}
			}
			if mine > 0 {
				board[p.x][p.y] = byte('0')+mine
			} else {
				board[p.x][p.y] = 'B'
				for i := range pointers {
					newSet[pointers[i]] = false
				}
			}
		}
		set = newSet
	}
	return board
}
