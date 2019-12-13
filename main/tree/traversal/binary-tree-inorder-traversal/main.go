package main

import (
	"LeetCodeGo/base"
	"LeetCodeGo/utils"
	"log"
)

/*

给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/


func traversal(node *base.TreeNode, nums *[]int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		traversal(node.Left, nums)
	}
	*nums = append(*nums, node.Val)
	if node.Right != nil {
		traversal(node.Right, nums)
	}
}

// 递归实现
func inorderTraversal(root *base.TreeNode) []int {
	nums := make([]int, 0)
	traversal(root, &nums)
	return nums
}

// 迭代实现
func inorderTraversal2(root *base.TreeNode) []int {
	var nums []int
	if root == nil {
		return nums
	}
	nodes := make([]*base.TreeNode, 0)
	cur := root
	for cur != nil || len(nodes) > 0 {
		for cur != nil {
			nodes = append(nodes, cur)
			cur = cur.Left
		}
		if len(nodes) > 0 {
			cur = nodes[len(nodes) - 1]  // 勿粗心大意，这里的cur不要使用 := 而使得go以为是局部变量
			nums = append(nums, cur.Val)
			nodes = nodes[:len(nodes) - 1]
			cur = cur.Right
		}
	}
	return nums
}

func main() {
	values := []string{"1", "nil", "2", "3"}
	root := base.GetBinaryTree(values)
	nums := inorderTraversal2(root)
	if len(nums) != 3 || !utils.CompareEqual(nums, []int{1, 3, 2}) {
		log.Fatal("error: ", nums)
	}
	log.Println("success")
}
