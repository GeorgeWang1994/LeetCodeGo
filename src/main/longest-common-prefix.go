package main

import (
	"log"
	"math"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	var minLen int = math.MaxInt32
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}

	if minLen == 0 {
		return ""
	}

	// 记录下最短的长度，每个字符串中的第idx个字符拿出来一一比较，
	var idx, maxIdx int
	for idx < minLen {
		var charArray []byte

		for _, str := range strs {
			charArray = append(charArray, str[idx])
		}
		firstChar := charArray[0]
		for _, char := range charArray[1:] {
			if firstChar != char {
				goto RETURN
			}
		}
		maxIdx += 1
		idx += 1
	}

	RETURN: return strs[0][:maxIdx]
}

func main() {
	result := longestCommonPrefix([]string {"dog","racecar","car"})
	if result != ""  {
		log.Printf("error: %v", result)
		return
	}

	result = longestCommonPrefix([]string {"flower","flow","flight"})
	if result != "fl"  {
		log.Printf("error: %v", result)
		return
	}

	log.Printf("success")
}
