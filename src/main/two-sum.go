package main

import (
	"fmt"
	"main/utils"
)

/*
	给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

	你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

	示例:

	给定 nums = [2, 7, 11, 15], target = 9

	因为 nums[0] + nums[1] = 2 + 7 = 9
	所以返回 [0, 1]

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/two-sum
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func twoSum(nums []int, target int) []int {
	// 遍历数组获取每个数字对应的差
	keySubList := make([]int, len(nums))
	for idx, num := range nums {
		keySubList[idx] = target - num
	}

	// 从map中进行匹配结果
	for idx, num := range nums {
		subNum := target - num
		if subNum == keySubList[idx] {
			firstN := 1
			for ; ;  {
				// 一直遍历，直到找到和当前下标不一致的下标，避免重复
				otherIdx := utils.IndexOf(nums, subNum, firstN)
				if otherIdx ==  -1 {
					break
				}
				if idx != otherIdx {
					return []int {idx, otherIdx}
				} else {
					firstN += 1
				}
			}
		}
	}
	return []int {}
}


func twoSum2(nums []int, target int) []int {
	/*
		利用map记录下当前的值对应的下标，重点在于一边遍历一遍设置
		因为如果提前设置好所有的下标，就会导致获取的下标不正确
		比如[3, 2, 4]，当出现差和当前值一样的情况，获取到两个数的下标就是重复的
	 */
	subMap := make(map[int]int)
	for idx, val := range nums {
		if subIdx, exist := subMap[target - val]; exist {
			return []int {idx, subIdx}
		} else {
			subMap[val] = idx
		}
	}
	return []int {}
}


func main() {
	data := []int{2, 7, 11, 15}
	result := twoSum(data, 9)

	if len(result) != 2 || !utils.CompareEqual(result, []int {0, 1}) {
		fmt.Printf("fail 1: %v", result)
		return
	}

	data = []int{3, 3, 11, 15}
	result = twoSum(data, 6)

	if len(result) != 2 || !utils.CompareEqual(result, []int {0, 1}) {
		fmt.Printf("fail 2: %v", result)
		return
	}

	data = []int{3, 2, 4}
	result = twoSum(data, 6)

	if len(result) != 2 || !utils.CompareEqual(result, []int {1, 2}) {
		fmt.Printf("fail 3: %v", result)
		return
	}

	fmt.Printf("success")
}
