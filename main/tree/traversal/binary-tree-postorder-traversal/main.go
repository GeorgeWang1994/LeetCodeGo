package main

import (
    "LeetCodeGo/base"
    "LeetCodeGo/utils"
    "container/list"
    "log"
)

/*
Given a binary tree, return the postorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [3,2,1]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/

func traversal(node *base.TreeNode, nums *[]int) {
    if node == nil {
        return
    }
    if node.Left != nil {
        traversal(node.Left, nums)
    }
    if node.Right != nil {
        traversal(node.Right, nums)
    }
    *nums = append(*nums, node.Val)
}

func postorderTraversal(root *base.TreeNode) []int {
    // 递归实现
    nums := make([]int, 0)
    traversal(root, &nums)
    return nums
}

/*

后序遍历的难点在于：需要判断上次访问的节点是位于左子树，还是右子树。

* 若是位于左子树，则需跳过根节点，先进入右子树，再回头访问根节点；
* 若是位于右子树，则直接访问根节点。

*/
func postorderTraversal2(root *base.TreeNode) []int {
    var nums []int
    if root == nil {
        return nums
    }
    nodes := make([]*base.TreeNode, 0)
    var prev *base.TreeNode
    cur := root

    for cur != nil || len(nodes) > 0 {
        // 不断将左子结点入栈
        for cur != nil {
            nodes = append(nodes, cur)
            cur = cur.Left
        }
        // 到这里说明当前结点为空，已经到了最左子结点，这时候需要出栈
        if len(nodes) > 0 {
            right := nodes[len(nodes)-1].Right
            // 跟节点被访问的前提是没有右子树或者右子树已经被访问过
            if right == nil || prev == right { // right == nil 表示不存在右结点，prev == right 表示右结点已经被访问过
                prev = nodes[len(nodes)-1]
                nums = append(nums, prev.Val)
                nodes = nodes[:len(nodes)-1]
            } else {
                // 如果存在右结点，则将当前结点设为跟节点，继续进行新一轮的左子树的遍历
                cur = right
            }
        }
    }
    return nums
}

func postorderTraversal3(root *base.TreeNode) []int {
    if root == nil {
        return nil
    }
    var prev *base.TreeNode
    var traver []int
    stack := list.New()
    stack.PushBack(root)
    for {
        e := stack.Back()
        if e == nil {
            break
        }
        node, _ := e.Value.(*base.TreeNode)
        // 1. 不存在左右子结点
        // 2. 左孩子和右孩子都在根结点前面被访问
        // （由于在方法postorderTraversal2中的设定，左子结点肯定会被访问到，因此只需要判断右子结点。
        // 	但是在方法postorderTraversal3中，和前面优先将左结点入栈不同，这里会按照 跟结点->右结点->左结点的顺序入栈）
        if (node.Left == nil && node.Right == nil) || (prev != nil && (node.Left == prev || node.Right == prev)) {
            traver = append(traver, node.Val)
            stack.Remove(e)
            prev = node
        } else {
            // 先将右子结点入栈，后将左子结点入栈，保证左子结点、右子结点在栈弹出时按照顺序访问
            if node.Right != nil {
                stack.PushBack(node.Right)
            }
            if node.Left != nil {
                stack.PushBack(node.Left)
            }
        }
    }
    return traver
}

func main() {
    values := []string{"1", "nil", "2", "3"}
    root := base.GetBinaryTree(values)
    nums := postorderTraversal3(root)
    if len(nums) != 3 || !utils.CompareEqual(nums, []int{3, 2, 1}) {
        log.Fatal("error: ", nums)
    }
    log.Println("success")
}
