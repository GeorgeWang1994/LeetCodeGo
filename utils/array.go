package utils

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

// 比较数组是否相等
func CompareEqual(list1, list2 []int) bool {
    if len(list1) != len(list2) {
        return false
    }

    if list1 == nil || list2 == nil {
        return false
    }

    for idx, val := range list1 {
        if val != list2[idx] {
            return false
        }
    }

    return true
}

// 获取数组中出现的下标，可以指定第几次出现的值
func IndexOf(list []int, target int, firstN int) int {
    cnt := 1
    for idx, val := range list {
        if val == target {
            if cnt == firstN {
                return idx
            }
            cnt += 1
        }
    }
    return -1
}

// 反转字符串数组（会修改原数组）
func ReverseArray(list []string) []string {
    for from, to := 0, len(list)-1; from < to; from, to = from+1, to-1 {
        list[from], list[to] = list[to], list[from]
    }

    return list
}

// 删除数组中重复的元素
func RemoveDuplicateArray(nums []int) []int {
    if len(nums) == 0 {
        return nums
    }

    result := make([]int, 0, len(nums))
    numsMap := make(map[int]struct{})
    for _, num := range nums {
        if _, ok := numsMap[num]; !ok {
            result = append(result, num)
            numsMap[num] = struct{}{}
        }
    }

    return result
}

// 删除数组中重复的元素，并且修改原来的数组
func RemoveDuplicateArrayWithChange(nums *[]int) {
    if len(*nums) == 0 {
        return
    }

    numsMap := make(map[int]struct{})
    lastLen := len(*nums)
    for i := 0; i < lastLen; i += 1 {
        num := (*nums)[i]
        if _, ok := numsMap[num]; !ok {
            numsMap[num] = struct{}{}
        } else {
            // 将后面的值全部改掉
            for j := i; j < lastLen-1; j += 1 {
                (*nums)[j] = (*nums)[j+1]
            }
            lastLen -= 1
            i -= 1
        }
    }
    *nums = (*nums)[:lastLen]
}
