package main

import (
	"LeetCodeGo/base"
	"LeetCodeGo/utils"
	"fmt"
	"log"
)

/*

根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func recursive(inorder []int, postorder []int) *base.TreeNode {
	if len(inorder) == 0 {
		return nil
	} else if len(inorder) == 1 {
		return &base.TreeNode{inorder[0], nil, nil}
	}

	rootVal := postorder[len(postorder)-1]
	root := &base.TreeNode{rootVal, nil, nil}
	rootIdx := utils.IndexOf(inorder, rootVal, 1)
	root.Left = recursive(inorder[:rootIdx], postorder[:rootIdx])
	root.Right = recursive(inorder[rootIdx+1:], postorder[rootIdx:len(postorder)-1])
	return root
}

func buildTree(inorder []int, postorder []int) *base.TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	return recursive(inorder, postorder)
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	result := buildTree(inorder, postorder)
	fmt.Println(result)
	log.Println("success")
}
