package main

import (
	"log"
	"strconv"
)

/*
	判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

	示例 1:

	输入: 121
	输出: true
	示例 2:

	输入: -121
	输出: false
	解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
	示例 3:

	输入: 10
	输出: false
	解释: 从右向左读, 为 01 。因此它不是一个回文数。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/palindrome-number
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

 */


// 判断数字是否是回文数
func isPalindrome(x int) bool {
	// 转为字符串再进行一个单词一个单词的比较
	xStr := strconv.Itoa(x)
	xLen := len(xStr)

	if xLen == 0 || xLen == 1 {
		return true
	}

	for from, to := 0, xLen - 1; from < to; from, to = from + 1, to - 1 {
		if xStr[from] != xStr[to] {
			return false
		}
	}

	return true
}


// 判断数字是否是回文数
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	// 一个数字一个数字的进行比较
	var from, to int = 0, x
	for x > 0 {
		from = from * 10 + x % 10
		x /= 10
	}

	return from == to
}


func main() {

	result := isPalindrome2(121)
	if !result {
		log.Printf("error: %v", result)
		return
	}

	result2 := isPalindrome2(-121)
	if result2 {
		log.Printf("error: %v", result2)
		return
	}

	result3 := isPalindrome2(70)
	if result3 {
		log.Printf("error: %v", result3)
		return
	}

	log.Printf("success")
}

