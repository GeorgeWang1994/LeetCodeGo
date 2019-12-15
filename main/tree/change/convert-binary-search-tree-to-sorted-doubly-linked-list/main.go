package main

import (
    "LeetCodeGo/base"
    "log"
)

func ConvertNode(root *base.TreeNode, last **base.TreeNode) {
    if root == nil {
        return
    }

    if root.Left != nil {
        ConvertNode(root.Left, last)
    }
    root.Left = *last
    if *last != nil {
        (*last).Right = root
    }
    *last = root
    if root.Right != nil {
        ConvertNode(root.Right, last)
    }
}

func Convert(root *base.TreeNode) *base.TreeNode {
    if root == nil {
        return nil
    }

    var lastNode *base.TreeNode
    ConvertNode(root, &lastNode)
    // 递归后lastNode指向树的最右结点，也是链表最后的结点，因此需要回到起始结点
    head := lastNode
    for head != nil && head.Left != nil {
        head = head.Left
    }
    return head
}

func main() {
    values := []string{"1", "2", "5", "3", "4", "nil", "6"}
    root := base.GetBinaryTree(values)
    result := Convert(root)
    if result.Val != 3 || result.Right.Val != 2 || result.Right.Left.Val != 3 {
        log.Fatal("error")
    }
    log.Print("success")
}
