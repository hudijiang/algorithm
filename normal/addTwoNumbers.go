package normal

import (
	"algorithm/base"
)
/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
示例：
	输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
	输出：7 -> 0 -> 8
	原因：342 + 465 = 807
*/
func addTwoNumbers(l1 *base.ListNode,l2 *base.ListNode) *base.ListNode{
	if l1==nil || l2==nil {
		return nil
	}
	// 初始化 指向第一个元素
	p:=l1.Next
	q:=l2.Next
	//进位标识符 如果进位则变更为1
	carry:=0

	var resultNode = new(base.ListNode)
	currNode := resultNode
	for  {
		//如果都为空则停止终止遍历
		if p==nil && q==nil {
			break
		}
		x:=q.Val
		if q==nil {
			x=0
		}
		y:=p.Val
		if p==nil {
			y=0
		}
		sum := x+y+carry
		carry=sum/10
		val:=sum%10
		node:=base.ListNode{val,nil}
		currNode.Next=&node
		currNode=currNode.Next
		if p!=nil {
			p=p.Next
		}
		if q!=nil {
			q=q.Next
		}
	}
	if carry != 0 {
		resultNode.Next = &base.ListNode{Val: carry}
	}
	return resultNode.Next
}


func addTwoNumbers2(l1 *base.ListNode, l2 *base.ListNode) *base.ListNode{
	var l *base.ListNode = &base.ListNode{}
	pre := l
	flag := 0
	for l1 != nil || l2 != nil {
		pre.Next = &base.ListNode{}
		p := pre.Next
		x := 0
		y := 0
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}
		p.Val = (x + y + flag) % 10
		flag = (x + y + flag) / 10
		pre = p
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if flag != 0 {
		pre.Next = &base.ListNode{Val: flag}
	}
	return l.Next
}



