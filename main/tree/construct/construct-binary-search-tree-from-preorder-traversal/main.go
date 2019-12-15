package main

import (
    "LeetCodeGo/base"
    "container/list"
    "fmt"
)

/*

返回与给定先序遍历 preorder 相匹配的二叉搜索树（binary search tree）的根结点。

(回想一下，二叉搜索树是二叉树的一种，其每个节点都满足以下规则，对于 node.left 的任何后代，值总 < node.val，而 node.right 的任何后代，值总 > node.val。此外，先序遍历首先显示节点的值，然后遍历 node.left，接着遍历 node.right。）


示例：

输入：[8,5,1,7,10,12]
输出：[8,5,10,1,7,null,12]

 https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/03/08/1266.png

提示：

1 <= preorder.length <= 100
先序 preorder 中的值是不同的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-search-tree-from-preorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func bstFromPreorder(preorder []int) *base.TreeNode {
    if len(preorder) == 0 {
        return nil
    } else if len(preorder) == 1 {
        return &base.TreeNode{preorder[0], nil, nil}
    }
    root := &base.TreeNode{preorder[0], nil, nil}
    rightIdx := len(preorder)
    for idx, val := range preorder {
        print(idx, val)
        if val > preorder[0] {
            rightIdx = idx
            break
        }
    }
    root.Left = bstFromPreorder(preorder[1:rightIdx])
    root.Right = bstFromPreorder(preorder[rightIdx:])
    return root
}

func bstFromPreorder2(preorder []int) *base.TreeNode {
    n := len(preorder)
    if n == 0 {
        return nil
    }

    statck := list.New()
    root := &base.TreeNode{preorder[0], nil, nil}
    statck.PushBack(root)

    for i := 1; i < n; i++ {
        parent := statck.Back().Value.(*base.TreeNode)
        child := &base.TreeNode{preorder[i], nil, nil}
        // 如果栈顶元素比子结点小，则一直出栈知道栈顶的值比子结点的值要大
        for {
            if statck.Len() == 0 {
                break
            }
            back := statck.Back().Value.(*base.TreeNode)
            if back.Val < child.Val {
                parent = statck.Back().Value.(*base.TreeNode)
                statck.Remove(statck.Back())
            } else {
                break
            }
        }
        if parent.Val < child.Val {
            parent.Right = child
        } else {
            parent.Left = child
        }
        statck.PushBack(child)
    }
    return root
}

func main() {
    preorder := []int{8, 5, 1, 7, 10, 12}
    result := bstFromPreorder2(preorder)
    fmt.Println(result)
}
