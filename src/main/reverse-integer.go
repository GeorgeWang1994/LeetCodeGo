package main

import (
	"log"
	"main/utils"
	"math"
	"strconv"
	"strings"
)

/*
	给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

	示例 1:

	输入: 123
	输出: 321
	 示例 2:

	输入: -123
	输出: -321
	示例 3:

	输入: 120
	输出: 21
	注意:

	假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。



	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/reverse-integer
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func reverse(x int) int {
	// 转正数->转字符串->反转字符串->再转成数字
	var isNegative bool = false
	if x < 0 {
		x = int(math.Abs(float64(x)))
		isNegative = true
	}
	xStr := strconv.Itoa(x)
	revStr := utils.ReverseString(xStr)
	revNum, _ := strconv.Atoi(revStr)
	PPowerNum := math.Pow(2, 31)
	if isNegative {
		revNum = -revNum
		if -PPowerNum > float64(revNum) {
			return 0
		}
	} else {
		if float64(revNum) >= PPowerNum {
			return 0
		}
	}
	return revNum
}


func reverse2(x int) int {
	// 转正数->转字符串->转字符串数组->反转数组->转字符串->再转成数字
	var isNegative bool = false
	if x < 0 {
		x = int(math.Abs(float64(x)))
		isNegative = true
	}
	xStr := strconv.Itoa(x)
	xArray := strings.Split(xStr, "")
	utils.ReverseArray(xArray)
	revStr := strings.Join(xArray, "")
	revNum, _ := strconv.Atoi(revStr)
	PPowerNum := math.Pow(2, 31)
	if isNegative {
		revNum = -revNum
		if -PPowerNum > float64(revNum) {
			return 0
		}
	} else {
		if float64(revNum) >= PPowerNum {
			return 0
		}
	}
	return revNum
}


func main()  {
	result := reverse2(-10000)
	if result != -1 {
		log.Printf("error: %v", result)
		return
	}

	result2 := reverse2(123)
	if result2 != 321 {
		log.Printf("error2: %v", result2)
		return
	}

	result3 := reverse2(-123)
	if result3 != -321 {
		log.Printf("error3: %v", result3)
		return
	}

	result4 := reverse2(1534236469)
	if result4 != 0 {
		log.Printf("error4: %v", result4)
		return
	}

	log.Printf("success")
}
