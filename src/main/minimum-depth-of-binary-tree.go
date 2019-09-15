package main

import (
	"container/list"
	"log"
	"main/base"
	"main/math"
)

/**
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

示例:

给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回它的最小深度  2.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-depth-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

 */

// dfs
func minDepth(root *base.TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)
	if leftDepth == 0 || rightDepth == 0 {
		return 1 + leftDepth + rightDepth
	}
	return 1 + math.Min(leftDepth, rightDepth)
}


func minDepth2(root *base.TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		// 此时才是叶子结点，也是题目所要求的最小深度的终止的位置，比如[1,2]的结果应该是2而不是1
		return 1
	}

	if root.Left != nil && root.Right != nil {
		return math.Min(minDepth(root.Left) + 1, minDepth(root.Right) + 1)
	} else if root.Left != nil {
		return minDepth(root.Left) + 1
	} else {
		return minDepth(root.Right) + 1
	}
}

// 使用自己实现的队列
func minDepth3(root *base.TreeNode) int {
	if root == nil {
		return 0
	}

	queue, err := base.NewNormalQueue(10)
	if err != nil {
		return 0
	}

	queue.Enqueue(root)
	minLevel := 0

	// 每遍历一次，相当于遍历一层的结点
	for queue.Length() > 0 {
		minLevel++
		levelSize := queue.Length()
		for i := 0; i < levelSize; i++ {
			node := queue.Front().Value().(*base.TreeNode)
			queue.Dequeue()
			// 只要出现该层出现第一个叶子结点，就是题目所要求的最短路径
			if node.Left == nil && node.Right == nil {
				return minLevel
			} else if node.Left != nil {
				queue.Enqueue(node.Left)
			} else {
				queue.Enqueue(node.Right)
			}
		}
	}
	return minLevel
}

// 使用原生的List
func minDepth4(root *base.TreeNode) int {
	var minLevel int
	queue := list.New()
	if root != nil {
		queue.PushBack(root)
	}
	for queue.Len() > 0 {
		minLevel++
		levelSize := queue.Len()
		for i := 0; i < levelSize; i++ {
			node := queue.Front().Value.(*base.TreeNode)
			queue.Remove(queue.Front())
			if node.Left == nil && node.Right == nil {
				return minLevel
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return minLevel
}


func main() {
	values := []string{"1", "2", "3", "nil", "4", "5", "nil"}
	root := base.GetBinaryTree(values)
	if minDepth2(root) != 3 {
		log.Fatal("error")
	}
	if minDepth3(root) != 3 {
		log.Fatal("error")
	}
	if minDepth4(root) != 3 {
		log.Fatal("error")
	}
	log.Print("success")
}
