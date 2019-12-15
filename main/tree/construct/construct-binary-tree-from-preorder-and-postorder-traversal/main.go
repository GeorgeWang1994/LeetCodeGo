package main

import (
    "LeetCodeGo/base"
    "LeetCodeGo/utils"
)

/*

返回与给定的前序和后序遍历匹配的任何二叉树。

 pre 和 post 遍历中的值是不同的正整数。

 

示例：

输入：pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
输出：[1,2,3,4,5,6,7]
 

提示：

1 <= pre.length == post.length <= 30
pre[] 和 post[] 都是 1, 2, ..., pre.length 的排列
每个输入保证至少有一个答案。如果有多个答案，可以返回其中一个。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func recursive(preorder []int, postorder []int) *base.TreeNode {
    if len(preorder) == 0 {
        return nil
    } else if len(preorder) == 1 {
        return &base.TreeNode{preorder[0], nil, nil}
    }

    root := &base.TreeNode{preorder[0], nil, nil}
    leftRootIdx := utils.IndexOf(postorder, preorder[1], 1)
    root.Left = recursive(preorder[1:1+leftRootIdx], postorder[:leftRootIdx])
    root.Right = recursive(preorder[1+leftRootIdx:], postorder[leftRootIdx+1:len(postorder)-1])
    return root
}

func constructFromPrePost(preorder []int, postorder []int) *base.TreeNode {
    if len(preorder) == 0 || len(postorder) == 0 {
        return nil
    }

    return recursive(preorder, postorder)
}
