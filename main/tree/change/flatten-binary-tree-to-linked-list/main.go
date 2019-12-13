package main

import (
	"LeetCodeGo/base"
	"log"
)

func flatten(root *base.TreeNode)  {
	if root == nil {
		return
	}

	for root != nil {
		left := root.Left
		right := root.Right
		if left != nil {
			// 找到根的左子树的最右子结点
			leftRight := left
			for leftRight.Right != nil  {
				leftRight = leftRight.Right
			}

			// 将根的右子结点连接到👆最右子结点的右结点（当前的右子树已经不存在了）
			leftRight.Right = right
			// 将根的左子树移动到右子树
			root.Right = left
			root.Left = nil
		}
		// 继续下一位
		root = root.Right
	}
}

func main() {
	values := []string{"1", "2", "5", "3", "4", "nil", "6"}
	root := base.GetBinaryTree(values)
	flatten(root)
	if root.Val != 1 {
		log.Fatal("error")
	}
	log.Print("success")
}
