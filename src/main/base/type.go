package base

import (
     "strconv"
)

type ListNode struct {
     Val int
     Next *ListNode
}


func GetListNode(nums []int) *ListNode {
     if len(nums) == 0 {
          return nil
     }

     var res = ListNode{Val:0}
     var ptr = &res
     for _, num := range nums {
          ptr.Next = &ListNode{Val:num}
          ptr = ptr.Next
     }

     return res.Next
}


type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}


func getTreeNodeByVal(value string) *TreeNode {
     if value == "nil" {
          return nil
     }
     valInt, err := strconv.Atoi(value)
     if err != nil {
          return nil
     }
     return &TreeNode{Val:valInt}
}


func GetBinaryTree(values []string) *TreeNode {
     length := len(values)
     if length == 0 {
          return nil
     }

     // 跟结点
     root := getTreeNodeByVal(values[0])
     treeNodes := make([]*TreeNode, 0, length * 2)
     treeNodes = append(treeNodes, root)

     idx := 0
     for idx <= length/2 {
          node := treeNodes[0]
          treeNodes = treeNodes[1:]
          if idx * 2 + 1 < length {
               node.Left = getTreeNodeByVal(values[idx * 2 + 1])
               if node.Left != nil {
                    treeNodes = append(treeNodes, node.Left)
               }
          }
          if idx * 2 + 2 < length {
               node.Right = getTreeNodeByVal(values[idx * 2 + 2])
               if node.Right != nil {
                    treeNodes = append(treeNodes, node.Right)
               }
          }
          idx += 1
     }

     return root
}