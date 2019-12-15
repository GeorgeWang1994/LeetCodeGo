package main

import (
    "LeetCodeGo/base"
    "LeetCodeGo/utils"
    "log"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**

将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func mergeTwoLists(l1 *base.ListNode, l2 *base.ListNode) *base.ListNode {
    if l1 == nil && l2 == nil {
        return nil
    } else if l1 == nil {
        return l2
    } else if l2 == nil {
        return l1
    }

    var newNode *base.ListNode

    if l1.Val <= l2.Val {
        newNode = &base.ListNode{Val: l1.Val, Next: nil}
        newNode.Next = mergeTwoLists(l1.Next, l2)
    } else {
        newNode = &base.ListNode{Val: l2.Val, Next: nil}
        newNode.Next = mergeTwoLists(l1, l2.Next)
    }

    return newNode
}

func mergeTwoLists2(l1 *base.ListNode, l2 *base.ListNode) *base.ListNode {
    if l1 == nil && l2 == nil {
        return nil
    } else if l1 == nil {
        return l2
    } else if l2 == nil {
        return l1
    }

    var root base.ListNode
    pre := &root
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            pre.Next = l1
            l1 = l1.Next
        } else {
            pre.Next = l2
            l2 = l2.Next
        }
        pre = pre.Next
    }
    if l1 != nil {
        pre.Next = l1
    }
    if l2 != nil {
        pre.Next = l2
    }

    return root.Next
}

func main() {
    l1Node := base.GetListNode([]int{})
    l2Node := base.GetListNode([]int{2, 4, 6, 8, 10})
    resNums := []int{2, 4, 6, 8, 10}
    result := mergeTwoLists(l1Node, l2Node)
    if result == nil || !utils.CheckListNodeWithListCorrect(result, resNums) {
        log.Printf("error: %v", result)
        return
    }

    l1Node = base.GetListNode([]int{1, 3, 5, 7, 9})
    l2Node = base.GetListNode([]int{2, 4, 6, 8, 10})
    resNums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    result = mergeTwoLists(l1Node, l2Node)
    if result == nil || !utils.CheckListNodeWithListCorrect(result, resNums) {
        log.Printf("error: %v", result)
        return
    }

    l1Node = base.GetListNode([]int{1, 3, 5, 7, 9})
    l2Node = base.GetListNode([]int{2, 4, 6, 8, 10})
    resNums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    result = mergeTwoLists2(l1Node, l2Node)
    if result == nil || !utils.CheckListNodeWithListCorrect(result, resNums) {
        log.Printf("error: %v", result)
        return
    }

    log.Printf("success")
}
