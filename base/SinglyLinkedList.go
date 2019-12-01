package base

import "fmt"

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

//create a new link list element 创建一个新的链表
func NewLinkList() *LinkList {
	return &LinkList{NewNode(0), 0}
}

//Append a Node element to an element 将元素添加到某个元素后面
func (list *LinkList) InsertAfter(element *Node, val interface{}) bool {
	if nil == element {
		return false
	}
	//创建一个新元素
	newNode := NewNode(val)
	//保存现有元素的下一个元素
	oldNode := element.next
	//将元素的下一个元素指向这个新的节点
	element.next = newNode
	//新元素的下一个元素指向原来的
	newNode.next = oldNode
	list.size++
	return true
}

//在某个节点前面插入节点
func (list *LinkList) InsertBefore(p *Node, val interface{}) bool {
	if nil == p || p == list.head { //如果指定的元素为空或者指定元素为头节点则直接返回
		return false
	}
	cur := list.head.next
	pre := list.head
	for nil != cur {
		if p == cur {
			break
		}
		pre = cur
		cur = cur.next
	}
	newNode := NewNode(val)
	pre.next = newNode
	newNode.next = cur
	list.size++
	return true
}

//在链表头部插入节点
func (list *LinkList) InsertToHead(node interface{}) bool {
	return list.InsertAfter(list.head, node)
}

//在链表尾部插入节点
func (list *LinkList) InsertToTail(val interface{}) bool {
	cur := list.head
	for nil != cur.next {
		cur = cur.next
	}
	return list.InsertAfter(cur, val)
}

//通过索引查找节点
func (list *LinkList) FindByIndex(index uint) *Node {
	if index > list.size {
		return nil
	}
	cur := list.head
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

//删除传入的节点
func (list *LinkList) DelNode(node *Node) bool {
	if nil == node {
		return false
	}
	cur := list.head.next
	pre := list.head
	for nil != cur {
		if cur == node {
			break
		}
		pre = cur
		cur = cur.next
	}
	pre.next = node.next
	node = nil
	list.size--
	return true
}

//打印链表
func (list *LinkList) Print() {
	cur := list.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetVal())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
