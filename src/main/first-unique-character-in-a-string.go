package main

import "log"

/**
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:

s = "leetcode"
返回 0.

s = "loveleetcode",
返回 2.
 

注意事项：您可以假定该字符串只包含小写字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-unique-character-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}

	charCntMap := make(map[int32]int)
	for _, c := range s {
		cnt, ok := charCntMap[c]
		if ok {
			charCntMap[c] = cnt + 1
		} else {
			charCntMap[c] = 1
		}
	}

	for idx, c := range s {
		if cnt, ok := charCntMap[c]; ok && cnt == 1 {
			return idx
		}
	}

	return -1
}


func firstUniqChar2(s string) int {
	counts := make([]int, 26)

	for _, item := range s {
		counts[item - 'a']++
	}

	for idx, item := range s {
		if counts[item - 'a'] == 1 {
			return idx
		}
	}

	return -1
}


func main() {
	result := firstUniqChar("leetcode")
	if result != 0 {
		log.Printf("error: %v", result)
		return
	}

	result = firstUniqChar("loveleetcode")
	if result != 2 {
		log.Printf("error: %v", result)
		return
	}

	result = firstUniqChar2("loveleetcode")
	if result != 2 {
		log.Printf("error: %v", result)
		return
	}

	log.Printf("success")
}
