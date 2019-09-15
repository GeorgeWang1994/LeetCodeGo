package base

import (
     "errors"
     "strconv"
)

type ListNode struct {
     Val int
     Next *ListNode
}


func GetListNode(nums []int) *ListNode {
     if len(nums) == 0 {
          return nil
     }

     var res = ListNode{Val:0}
     var ptr = &res
     for _, num := range nums {
          ptr.Next = &ListNode{Val:num}
          ptr = ptr.Next
     }

     return res.Next
}


type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}


func getTreeNodeByVal(value string) *TreeNode {
     if value == "nil" {
          return nil
     }
     valInt, err := strconv.Atoi(value)
     if err != nil {
          return nil
     }
     return &TreeNode{Val:valInt}
}


func GetBinaryTree(values []string) *TreeNode {
     length := len(values)
     if length == 0 {
          return nil
     }

     // 跟结点
     root := getTreeNodeByVal(values[0])
     treeNodes := make([]*TreeNode, 0, length * 2)
     treeNodes = append(treeNodes, root)

     idx := 0
     for idx <= length/2 {
          node := treeNodes[0]
          treeNodes = treeNodes[1:]
          if idx * 2 + 1 < length {
               node.Left = getTreeNodeByVal(values[idx * 2 + 1])
               if node.Left != nil {
                    treeNodes = append(treeNodes, node.Left)
               }
          }
          if idx * 2 + 2 < length {
               node.Right = getTreeNodeByVal(values[idx * 2 + 2])
               if node.Right != nil {
                    treeNodes = append(treeNodes, node.Right)
               }
          }
          idx += 1
     }

     return root
}


type LinkNode struct {
     value interface{}
     prev *LinkNode
     next *LinkNode
}

func (node *LinkNode) Value() interface{} {
     return node.value
}

func (node *LinkNode) Set(value interface{}) {
     node.value = value
}

func (node *LinkNode) Previous() *LinkNode {
     return node.prev
}

func (node *LinkNode) Next() *LinkNode {
     return node.next
}


type NormalQueue struct {
     front    *LinkNode
     rear     *LinkNode
     length   int
     capacity int
}

func NewNormalQueue(capacity int) (*NormalQueue, error) {
     if capacity <= 0 {
          return nil, errors.New("capacity is less than 0")
     }

     front := &LinkNode{
          value:    nil,
          prev: nil,
     }

     rear := &LinkNode{
          value:    nil,
          prev: front,
     }

     front.next = rear
     return &NormalQueue{
          front:    front,
          rear:     rear,
          capacity: capacity,
     }, nil
}

func (q *NormalQueue) Length() int {
     return q.length
}

func (q *NormalQueue) Capacity() int {
     return q.capacity
}

func (q *NormalQueue) Front() *LinkNode {
     if q.length == 0 {
          return nil
     }

     return q.front.next
}

func (q *NormalQueue) Rear() *LinkNode {
     if q.length == 0 {
          return nil
     }

     return q.rear.prev
}

func (q *NormalQueue) Enqueue(value interface{}) bool {
     if q.length == q.capacity || value == nil {
          return false
     }

     node := &LinkNode{
          value: value,
     }

     if q.length == 0 {
          q.front.next = node
     }

     node.prev = q.rear.prev
     node.next = q.rear
     q.rear.prev.next = node
     q.rear.prev = node
     q.length++

     return true
}

func (q *NormalQueue) Dequeue() interface{} {
     if q.length == 0 {
          return nil
     }

     result := q.front.next
     q.front.next = result.next
     result.next = nil
     result.prev = nil
     q.length--

     return result.value
}