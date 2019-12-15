package main

import (
    "LeetCodeGo/base"
    "fmt"
    "math"
)

func getDecimalValue(head *base.ListNode) int {
    if head == nil {
        return 0
    }

    // 链表反转
    var tail *base.ListNode
    cur, prev := head, tail
    for cur != nil {
        next := cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }

    head = prev
    idx, sum := 0, 0
    for head != nil {
        if head.Val != 0 {
            sum += int(math.Pow(float64(2), float64(idx)))
        }
        idx++
        head = head.Next
    }
    return sum
}

func main() {
    node := base.GetListNode([]int{1, 0, 1})
    result := getDecimalValue(node)
    fmt.Println(result)
}
