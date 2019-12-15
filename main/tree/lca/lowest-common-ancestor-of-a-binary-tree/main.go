package main

import (
    "LeetCodeGo/base"
    "log"
    "math"
)

/*

给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]



示例 1:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
示例 2:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
 

说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉树中。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

var ans *base.TreeNode

func containCounter(root, p, q *base.TreeNode) int {
    if root == nil {
        return 0
    }
    mid := 0
    if root == p || root == q {
        mid = 1
    }
    left := containCounter(root.Left, p, q)
    // 说明中间节点和左子结点存在p或者q
    if mid+left == 2 {
        if ans == nil {
            ans = root
        }
        return 2
    }
    right := containCounter(root.Right, p, q)
    // 说明左右结点存在p或者q，或者中间结点和右结点存在p或者q
    if mid+left+right == 2 {
        if ans == nil {
            ans = root
        }
    }
    return mid + left + right
}

func lowestCommonAncestor(root, p, q *base.TreeNode) *base.TreeNode {
    ans = nil // 重置结果，避免上次测试用例影响
    containCounter(root, p, q)
    return ans
}

// 计算跟节点到目标结点的路径
func findTargetPath(root, target *base.TreeNode) []*base.TreeNode {
    stack := make([]*base.TreeNode, 0)
    var lastVisit *base.TreeNode
    cur := root
    for cur != nil || len(stack) > 0 {
        if cur != nil {
            stack = append(stack, cur)
            cur = cur.Left
        } else {
            peek := stack[len(stack)-1]
            // 如果存在右结点，并且没有被访问过，则将当前结点设置为右结点
            if peek.Right != nil && lastVisit != peek.Right {
                cur = peek.Right
            } else {
                // 如果不存在右结点或者左右结点已经都访问过
                if peek == target {
                    return stack
                }
                lastVisit = peek
                stack = stack[:len(stack)-1]
                cur = nil
            }
        }
    }
    return stack
}

func lowestCommonAncestor2(root, p, q *base.TreeNode) *base.TreeNode {
    pathP, pathQ := findTargetPath(root, p), findTargetPath(root, q)
    lenP, lenQ := len(pathP), len(pathQ)
    var prev *base.TreeNode
    for i := 0; i < int(math.Min(float64(lenP), float64(lenQ))); i++ {
        if pathP[i] != pathQ[i] {
            return prev
        }
        prev = pathP[i]
    }
    return prev
}

func main() {
    values := []string{"3", "5", "1", "6", "2", "0", "8", "nil", "nil", "7", "4"}
    root := base.GetBinaryTree(values)
    if lowestCommonAncestor2(root, root.Left, root.Left.Right.Right).Val != 5 {
        log.Fatal("error")
    }
    log.Print("success")
}
