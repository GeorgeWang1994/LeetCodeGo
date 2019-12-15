package main

import (
    "LeetCodeGo/base"
    "log"
    "strconv"
    "strings"
)

func recursive(root *base.TreeNode, curPath []string, paths *[]string) {
    if root == nil {
        return
    }
    // 将当前结点加入到路径中（也可以初始值模版就添加跟节点，后面只需要加左右子结点即可）
    curPath = append(curPath, strconv.Itoa(root.Val))
    if root.Left == nil && root.Right == nil {
        *paths = append(*paths, strings.Join(curPath, "->"))
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

func binaryTreePaths(root *base.TreeNode) []string {
    curPath, paths := make([]string, 0), make([]string, 0)
    if root == nil {
        return paths
    }
    recursive(root, curPath, &paths)
    return paths
}

func main() {
    values := []string{"1", "2", "3", "nil", "5", "nil", "nil"}
    root := base.GetBinaryTree(values)
    result := binaryTreePaths(root)
    log.Printf("%v", result)
}
