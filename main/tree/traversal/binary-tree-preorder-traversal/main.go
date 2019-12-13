package main

import (
	"LeetCodeGo/base"
	"LeetCodeGo/utils"
	"container/list"
	"log"
)

/*
Given a binary tree, return the preorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,2,3]
Follow up: Recursive solution is trivial, could you do it iteratively?
 */

func traversal(node *base.TreeNode, nums *[]int) {
	if node == nil {
		return
	}
	*nums = append(*nums, node.Val)
	if node.Left != nil {
		traversal(node.Left, nums)
	}
	if node.Right != nil {
		traversal(node.Right, nums)
	}
}

// 递归实现
func preorderTraversal(root *base.TreeNode) []int {
	nums := make([]int, 0)
	traversal(root, &nums)
	return nums
}

// 迭代实现
func preorderTraversal2(root *base.TreeNode) []int {
	var nums []int
	if root == nil {
		return nums
	}
	nodes := make([]*base.TreeNode, 0)
	cur := root
	for cur != nil || len(nodes) > 0 {
		for cur != nil {
			nodes = append(nodes, cur)
			nums = append(nums, cur.Val)  // 上半部分处理左右子节点，下半部分处理跟节点
			cur = cur.Left
		}
		if len(nodes) > 0 {
			cur = nodes[len(nodes) - 1]  // 勿粗心大意，这里的cur不要使用 := 而使得go以为是局部变量
			nodes = nodes[:len(nodes) - 1]
			cur = cur.Right
		}
	}
	return nums
}

// 要保证根结点在左孩子和右孩子访问之后才能访问，因此对于任一结点P，先将其入栈。如果P不存在左孩子和右孩子，则可以直接访问它；
// 或者P存在左孩子或者右孩子，但是其左孩子和右孩子都已被访问过了，则同样可以直接访问该结点。
// 若非上述两种情况，则将P的右孩子和左孩子依次入栈，这样就保证了每次取栈顶元素的时候，左孩子在右孩子前面被访问，左孩子和右孩子都在根结点前面被访问。
func preorderTraversal3(root *base.TreeNode) []int {
	if root == nil {
		return nil
	}
	var traver []int
	stack := list.New()
	stack.PushBack(root)
	for {
		e := stack.Back()
		if e == nil {
			break
		}
		node, _ := e.Value.(*base.TreeNode)
		traver = append(traver, node.Val)
		stack.Remove(e)
		// 先将右子结点入栈，后将左子结点入栈，保证左子结点、右子结点在栈弹出时按照顺序访问
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return traver
}

func main() {
	values := []string{"1", "nil", "2", "3"}
	root := base.GetBinaryTree(values)
	nums := preorderTraversal3(root)
	if len(nums) != 3 || !utils.CompareEqual(nums, []int{1, 2, 3}) {
		log.Fatal("error: ", nums)
	}
	log.Println("success")
}
