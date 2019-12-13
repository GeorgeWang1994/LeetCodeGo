package main

import (
	"fmt"
	"math"
	"sort"
)

type IntList []int

func (i IntList) Len() int {
	return len(i)
}

func (i IntList) Swap(p, q int) {
	i[p], i[q] = i[q], i[p]
}

func (i IntList) Less(p, q int) bool {
	return i[p] < i[q]
}

/*

难点:
1. 需要考虑到时间复杂度
2. 从除数最小值到除数最大值，计算出来的除数之和在递减

 */
func smallestDivisor(nums []int, threshold int) int {
	l := len(nums)
	sort.Sort(IntList(nums))
	low, high := 1, nums[l-1]
	target := high  // 保证目标是数组中最大值，可保证除数和最小（因为此时全部除数都是1，除数之和就是 1 * l）
	for low <= high {
		mid := (low + high) / 2
		sum := 0
		for _, v := range nums {
			sum += int(math.Ceil(float64(v) / float64(mid)))  // 将结果向上取整
		}
		// 如果 mid 的结果比 threshold 要小，说明满足要求或者 mid 太大了，继续在 [low, mid-1] 中继续查找
		// 如果 mid 的结果比 threshold 要大，说明满足要求或者 mid 太小了，继续在 [mid+1, high] 中继续查找
		if sum <= threshold {
			target = int(math.Min(float64(target), float64(mid)))
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return target
}

func main()  {
	result := smallestDivisor([]int{1,2,5,9}, 6)
	fmt.Println(result)
}
