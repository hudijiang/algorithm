package normal

import (
	"algorithm/base"
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	listNode := base.ListNode{0,&base.ListNode{2,&base.ListNode{4,&base.ListNode{3,nil}}}}
	listNode1 := base.ListNode{0,&base.ListNode{5,&base.ListNode{6,&base.ListNode{4,nil}}}}
	nodeList:=addTwoNumbers(&listNode,&listNode1)
	curNode := nodeList
	for   {
		fmt.Println(curNode.Val)
		curNode=curNode.Next
		if curNode==nil {
			break
		}
	}

}



