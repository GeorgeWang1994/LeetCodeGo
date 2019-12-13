package main

import (
	"LeetCodeGo/base"
	"LeetCodeGo/utils"
	"container/list"
)

/*

给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。这个二叉树与满二叉树（full binary tree）结构相同，但一些节点为空。

每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。

示例 1:

输入:

           1
         /   \
        3     2
       / \     \
      5   3     9

输出: 4
解释: 最大值出现在树的第 3 层，宽度为 4 (5,3,null,9)。
示例 2:

输入:

          1
         /
        3
       / \
      5   3

输出: 2
解释: 最大值出现在树的第 3 层，宽度为 2 (5,3)。
示例 3:

输入:

          1
         / \
        3   2
       /
      5

输出: 2
解释: 最大值出现在树的第 2 层，宽度为 2 (3,2)。
示例 4:

输入:

          1
         / \
        3   2
       /     \
      5       9
     /         \
    6           7
输出: 8
解释: 最大值出现在树的第 4 层，宽度为 8 (6,null,null,null,null,null,null,7)。
注意: 答案在32位有符号整数的表示范围内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-width-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

type NodeInfo struct {
	node *base.TreeNode;
	idx  int;
}

// 问题可以转化为求每一层最左结点和最右结点的距离，取所有层的距离最大
func widthOfBinaryTree(root *base.TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	}
	var maxWidth = -1
	stack := list.New()
	stack.PushBack(NodeInfo{root, 0})
	tempStack := list.New()
	for {
		info := stack.Front().Value.(NodeInfo)
		stack.Remove(stack.Front())

		if info.node.Left != nil {
			tempStack.PushBack(NodeInfo{info.node.Left, info.idx * 2 + 1})
		}
		if info.node.Right != nil {
			tempStack.PushBack(NodeInfo{info.node.Right, info.idx * 2 + 2})
		}

		if stack.Len() == 0 {
			if tempStack.Len() == 0 {
				break
			}
			maxWidth = utils.Max(maxWidth, (tempStack.Back().Value.(NodeInfo).idx - tempStack.Front().Value.(NodeInfo).idx) + 1)
			stack.PushBackList(tempStack)
			tempStack = list.New()
		}
	}
	return maxWidth
}
