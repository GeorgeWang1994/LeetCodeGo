package main

import (
	"strconv"
	"strings"
)

func subtractProductAndSum(n int) int {
	nStr := strconv.Itoa(n)
	nStrList := strings.Split(nStr, "")
	multi, sum := 1, 0
	for _, v := range nStrList {
		vint, _ := strconv.Atoi(v)
		multi *= vint
		sum += vint
	}
	return multi - sum
}
