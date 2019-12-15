package utils

import "container/list"

type Stack struct {
    list *list.List
}

// 新建栈
func NewStack() *Stack {
    array := list.New()
    return &Stack{array}
}

// 插入数据到栈顶
func (stack *Stack) Push(value interface{}) {
    stack.list.PushBack(value)
}

// 从栈顶中弹出数据
func (stack *Stack) Pop() interface{} {
    back := stack.list.Back()
    if back != nil {
        stack.list.Remove(back)
        return back.Value
    }
    return nil
}

// 从栈顶中获取数据
func (stack *Stack) Peak() interface{} {
    back := stack.list.Back()
    if back != nil {
        return back.Value
    }
    return nil
}

// 栈长度
func (stack *Stack) Len() interface{} {
    return stack.list.Len()
}

// 栈是否为空
func (stack *Stack) Empty() interface{} {
    return stack.Len() == 0
}
