package base

import (
	"LeetCodeGo/base"
	"container/list"
)

// 用栈迭代实现获取从跟结点到叶子结点的路径
func GetPathStackIterate(root *base.TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	curPath := make([]int, 0)
	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		peek := stack.Back().Value.(*base.TreeNode)
		curPath = append(curPath, peek.Val)
		stack.Remove(stack.Back())

		if peek.Left == nil && peek.Right == nil {
			result = append(result, curPath)
		} else {
			if peek.Left != nil {
				stack.PushBack(peek.Left)
				curPath = append(curPath, peek.Left.Val)
			}
			if peek.Right != nil {
				stack.PushBack(peek.Right)
				curPath = append(curPath, peek.Right.Val)
			}
		}
		curPath = curPath[:len(curPath)-1]
	}
	return result
}

// 用DFS递归实现获取从跟结点到叶子结点的路径
func recursive(root *base.TreeNode, curPath []int, paths *[][]int) {
	if root == nil {
		return
	}
	// 将当前结点加入到路径中（也可以初始值模版就添加跟节点，后面只需要加左右子结点即可）
	curPath = append(curPath, root.Val)
	if root.Left == nil && root.Right == nil {
		*paths = append(*paths, curPath)
		return
	}
	if root.Left != nil {
		recursive(root.Left, curPath, paths)
	}
	if root.Right != nil {
		recursive(root.Right, curPath, paths)
	}
	// 记得回退，也可以在调用函数的curPath里加上对应结点的值(Python可以)，这样就可以不用手动回退了
	curPath = curPath[:len(curPath)-1]
}

func GetPathDFSRecursive(root *base.TreeNode) [][]int {
	curPath, paths := make([]int, 0), make([][]int, 0)
	if root == nil {
		return paths
	}
	recursive(root, curPath, &paths)
	return paths
}