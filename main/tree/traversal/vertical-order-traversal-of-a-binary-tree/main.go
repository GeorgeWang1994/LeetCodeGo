package main

import (
    "LeetCodeGo/base"
    "LeetCodeGo/utils"
    "fmt"
    "sort"
)

type NodeInfo struct {
    value int
    y     int
}

type NodeInfoSlice []NodeInfo

func (slice NodeInfoSlice) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}

func (slice NodeInfoSlice) Len() int {
    return len(slice)
}

// 按照y坐标轴从大到小排序，如果相等则按照值从小到大排序
func (slice NodeInfoSlice) Less(i, j int) bool {
    if slice[i].y > slice[j].y {
        return true
    } else if slice[i].y == slice[j].y && slice[i].value < slice[j].value {
        return true
    } else {
        return false
    }
}

func recursive(node *base.TreeNode, x, y int, memory *map[int][]NodeInfo) {
    if node == nil {
        return
    }

    info := NodeInfo{value: node.Val, y: y}
    if ys, ok := (*memory)[x]; ok {
        ys := append(ys, info)
        (*memory)[x] = ys
    } else {
        ys := []NodeInfo{info}
        (*memory)[x] = ys
    }

    if node.Left != nil {
        recursive(node.Left, x-1, y-1, memory)
    }

    if node.Right != nil {
        recursive(node.Right, x+1, y-1, memory)
    }
}

func verticalTraversal(root *base.TreeNode) [][]int {
    result := make([][]int, 0)
    if root == nil {
        return result
    }

    memory := make(map[int][]NodeInfo, 0)
    recursive(root, 0, 0, &memory)

    keys := make([]int, 0)
    for k := range memory {
        keys = append(keys, k)
    }

    sort.Sort(utils.IntList(keys))

    for _, key := range keys {
        values := memory[key]
        sort.Sort(NodeInfoSlice(values))
        ys := make([]int, 0)
        for _, v := range values {
            ys = append(ys, v.value)
        }
        result = append(result, ys)
    }

    return result
}

func main() {
    values := []string{"0", "2", "1", "3", "nil", "nil", "nil", "4", "5", "nil", "7", "6", "nil", "10", "8", "11", "9"}
    root := base.GetBinaryTree(values)
    result := verticalTraversal(root)
    fmt.Println(result)
    values = []string{"1", "2", "3", "4", "5", "6", "7"}
    root = base.GetBinaryTree(values)
    result = verticalTraversal(root)
    fmt.Println(result)
}
