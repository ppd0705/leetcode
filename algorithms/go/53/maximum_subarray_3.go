package main

type status struct {
	lSum, mSum, rSUm, iSum int
}

func maxSubArray3(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}

func pushUp(l, r status) status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSUm, r.iSum+l.rSUm)
	mSum := max(max(l.mSum, r.mSum), l.rSUm+r.lSum)
	return status{lSum, mSum, rSum, iSum}
}

func get(nums []int, l, r int) status {
	if l == r {
		return status{nums[l], nums[l], nums[l], nums[l]}
	}

	m := (l + r) / 2
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}
