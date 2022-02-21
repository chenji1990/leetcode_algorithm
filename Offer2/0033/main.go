package main

func verifyPostorder(postorder []int) bool {
	var recur func(start, end int) bool
	recur = func(start, end int) bool {
		if start >= end {
			return true
		}
		head := postorder[end]
		leftEnd := start
		for postorder[leftEnd] < head {
			leftEnd += 1
		}
		rightEnd := leftEnd
		for postorder[rightEnd] > head {
			rightEnd += 1
		}
		return rightEnd == end && recur(start, leftEnd-1) && recur(leftEnd, rightEnd-1)
	}
	return recur(0, len(postorder)-1)
}
