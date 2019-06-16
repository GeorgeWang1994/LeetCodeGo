package utils

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
	for from, to := 0, len(list)-1; from < to; from, to = from + 1, to - 1 {
		list[from], list[to] = list[to], list[from]
	}

	return list
}

