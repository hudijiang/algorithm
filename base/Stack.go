package base

// 定义一个栈，栈数据结构特点：后进先出
type Stack struct {
	el []interface{}
}

//this stack add a element
func (stack *Stack) Push(el interface{}){
	// 创建一个新的切片并将现有的数据加入到新切片的尾部作为该栈
	stack.el=append([]interface{}{el},stack.el[0:]...)
}

//remove this stack frist element，return frist elements
func (stack *Stack) Pop() interface{}{
	el:=stack.el[0]
	stack.el=stack.el[1:]
	return el
}

//get this stack is empty,if this stack leng is zore return true, else return this stack false
func (stack *Stack) IsEmpty() bool {
	return len(stack.el)==0
}

//get this stack size
func (stack *Stack) Size() int {
	return len(stack.el)
}

