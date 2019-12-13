package utils

import (
	"LeetCodeGo/base"
)


// 检查链表是否和数组一致
func CheckListNodeWithListCorrect(root *base.ListNode, nums []int) bool {
	if root == nil && len(nums) != 0 {
		return false
	} else if root != nil && len(nums) == 0 {
		return false
	}

	var p = root
	for _, num := range nums {
		if p == nil {
			return false
		}

		if p.Val != num {
			return false
		}
		p = p.Next
	}

	return true
}
