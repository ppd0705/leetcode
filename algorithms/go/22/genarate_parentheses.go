package main

func genarate(left int, right int, patten string, res *[]string) {
	if right == 0 {
		*res = append(*res, patten)
		return
	}
	if left > 0 {
		genarate(left-1, right, patten+"(", res)
	}
	if right > left {
		genarate(left, right-1, patten+")", res)
	}
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	if n > 0 {
		genarate(n, n, "", &res)
	}
	return res
}
