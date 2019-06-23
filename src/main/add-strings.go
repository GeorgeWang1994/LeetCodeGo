package main

import (
	"log"
	"strconv"
)

/**

给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

注意：

num1 和num2 的长度都小于 5100.
num1 和num2 都只包含数字 0-9.
num1 和num2 都不包含任何前导零。
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

 */


func addStrings(num1 string, num2 string) string {
	decimal := 0  // 进位
	len1, len2 := len(num1), len(num2)
	i, j := len1 - 1, len2 - 1
	result := ""

	for ; i >= 0 || j >= 0; i, j = i - 1, j - 1  {
		cur1, cur2 := 0, 0
		if i >= 0 {
			cur1 = int(num1[i] - '0')
		}
		if j >= 0 {
			cur2 = int(num2[j] - '0')
		}

		cur := cur1 + cur2

		if decimal != 0 {
			cur += decimal
		}

		if cur >= 10 {
			decimal = cur / 10
		} else {
			decimal = 0
		}
		result = strconv.Itoa(cur % 10) + result
	}

	if decimal != 0 {
		result = "1" + result
	}

	return result
}


func addStrings2(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	len1, len2 := len(num1), len(num2)
	result := make([]byte, len1)

	sum := byte(0)
	for i, j := len1 - 1, len2 - 1; i >= 0 || j >= 0; i, j, sum = i - 1, j - 1, sum/10  {
		if j >= 0 {
			sum += num2[j] - '0'
		}
		sum += num1[i] - '0'
		result[i] = sum % 10 + '0'
	}

	if sum != 0 {
		result = append([]byte{'1'}, result...)
	}

	return string(result)
}


func main() {
	result := addStrings("1", "1")
	if result != "2" {
		log.Printf("error: %v", result)
		return
	}

	result = addStrings("4999", "1")
	if result != "5000" {
		log.Printf("error: %v", result)
		return
	}

	result = addStrings("999", "1")
	if result != "1000" {
		log.Printf("error: %v", result)
		return
	}

	result = addStrings2("999", "1")
	if result != "1000" {
		log.Printf("error: %v", result)
		return
	}

	result = addStrings2("99", "9")
	if result != "108" {
		log.Printf("error: %v", result)
		return
	}

	log.Printf("success")
}
