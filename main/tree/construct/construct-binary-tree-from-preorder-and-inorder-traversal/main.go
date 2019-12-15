package main

import (
    "LeetCodeGo/base"
    "LeetCodeGo/utils"
    "fmt"
    "log"
)

/*

根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

*/

func recursive(preorder []int, inorder []int) *base.TreeNode {
    if len(inorder) == 0 {
        return nil
    } else if len(inorder) == 1 {
        return &base.TreeNode{inorder[0], nil, nil}
    }

    root := &base.TreeNode{preorder[0], nil, nil}
    rootIdx := utils.IndexOf(inorder, preorder[0], 1)
    root.Left = recursive(preorder[1:1+rootIdx], inorder[:rootIdx])
    root.Right = recursive(preorder[1+rootIdx:], inorder[rootIdx+1:])
    return root
}

func buildTree(preorder []int, inorder []int) *base.TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }

    return recursive(preorder, inorder)
}

func main() {
    //values := []string{"3", "9", "20", "nil", "nil", "15", "7"}
    preorder := []int{3, 9, 20, 15, 7}
    inorder := []int{9, 3, 15, 20, 7}

    result := buildTree(preorder, inorder)
    fmt.Println(result)
    log.Println("success")
}
