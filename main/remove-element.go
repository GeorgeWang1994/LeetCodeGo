package main

import "log"

/**
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-strstr
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


// 字符串
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	aLen, bLen := len(haystack), len(needle)
	if (aLen == 0) != (bLen == 0) {
		return -1
	}

	idx := 0
	for i := 0; i < len(haystack); i += 1 {
		if haystack[i] == needle[idx] {
			p, q := i, idx
			for ; p < len(haystack) && q < len(needle); p, q = p + 1, q + 1 {
				if haystack[p] != needle[q] {
					idx = 0
					break
				}
			}
			if q != len(needle) {
				idx = 0
				continue
			} else {
				return i
			}
		}
	}
	return -1
}


func main() {
	result := strStr("hello", "ll")
	if result != 2 {
		log.Printf("error: %v", result)
		return
	}

	result = strStr("aaaaa", "bba")
	if result != -1 {
		log.Printf("error: %v", result)
		return
	}

	result = strStr("", "")
	if result != 0 {
		log.Printf("error: %v", result)
		return
	}

	result = strStr("abc", "")
	if result != 0 {
		log.Printf("error: %v", result)
		return
	}

	log.Printf("success")
}
