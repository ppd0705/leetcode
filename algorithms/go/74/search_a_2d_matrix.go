package main

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}
	column := len(matrix[0])
	if column == 0 {
		return false
	}

	for _, m := range matrix {
		if m[column-1] >= target {
			l, r := 0, column-1
			for l <= r {
				mid := (l + r) / 2
				if m[mid] == target {
					return true
				} else if m[mid] > target {
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
			return false
		}
	}
	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}
	column := len(matrix[0])
	if column == 0 {
		return false
	}

	i, j := 0, row-1
	for i <= j {
		mid := (i + j) / 1
		if matrix[mid][column-1] == target {
			return true
		} else if matrix[mid][column-1] < target {
			i = mid + 1
		} else if matrix[mid][0] <= target {
			m := matrix[mid]
			l, r := 0, column-1
			for l <= r {
				mid := (l + r) / 2
				if m[mid] == target {
					return true
				} else if m[mid] > target {
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
			return false
		} else {
			j = mid - 1
		}
	}
	return false
}
