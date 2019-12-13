package main

import "fmt"

func groupThePeople(groupSizes []int) [][]int {
	userInfo := make(map[int][]int, 0)
	result := make([][]int, 0)

	for idx, val := range groupSizes {
		if users, ok := userInfo[val]; ok {
			users = append(users, idx)
			userInfo[val] = users
		} else {
			userInfo[val] = []int{idx}
		}
	}

	for cnt, users := range userInfo {
		if len(users) != cnt {
			temp := make([]int, 0)
			for _, v := range users {
				temp = append(temp, v)
				if len(temp) == cnt {
					result = append(result, temp)
					temp = make([]int, 0)
				}
			}
		} else {
			result = append(result, users)
		}
	}

	return result
}


func main() {
	result := groupThePeople([]int{3,3,3,3,3,1,3})
	fmt.Printf("%v", result)
}