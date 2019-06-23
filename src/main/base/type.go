package base

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
