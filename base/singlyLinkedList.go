package base

//declare a Node struct 声明一个Node结构体
type Node struct {
	next  *Node
	value interface{}
}

//declare a LinkList struct 声明一个链表结构体
type LinkList struct {
	head *Node
	size uint
}

//create a new Node element 创建一个新的节点并给定一个val
func NewNode(val interface{}) *Node {
	return &Node{nil, val}
}

//get this next Node element 获得当前 元素的下一个元素
func (node *Node) GetNext() *Node {
	return node.next
}

//get this Node value 获得当前元素的值
func (node *Node) GetVal() interface{} {
	return node.value
}
