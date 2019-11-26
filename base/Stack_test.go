package base

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
)
// stack test
func TestStack(t *testing.T) {
	stack := new(Stack)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(*stack)
	stack.Pop()
	fmt.Println(*stack)

	//stoage value stack
	valStack := new(Stack)
	//stoage ops stack
	opsStack := new(Stack)


	//模拟计算
	str:="((6-2)+(2*3)*3)+5"
	for _,ch:=range str {
		str:=string(ch)
		flag1,er1:=regexp.MatchString(`[0-9]`,str)
		flag2,er2:=regexp.MatchString(`[+|\-|*|/]`,str)
		if flag1&&er1==nil {
			num,error := strconv.ParseFloat(string(ch),64)
			if error==nil {
				valStack.Push(num)
			}
		} else if flag2&&er2==nil {
			opsStack.Push(str)
		} else if string(ch)==")" {
			num := valStack.Pop().(float64)
			var op = opsStack.Pop().(string)
			switch op {
			case "+":
				valStack.Push(num+valStack.Pop().(float64))
			case "-":
				valStack.Push(valStack.Pop().(float64)-num)
			case "*":
				valStack.Push(num*valStack.Pop().(float64))
			case "/":
				valStack.Push(valStack.Pop().(float64)/num)
			}
		}
	}
	fmt.Println("值栈为",*valStack)
	fmt.Println("操作栈为",*opsStack)
	for  {
		num:=valStack.Pop().(float64)
		var op = opsStack.Pop().(string)
		switch op {
		case "+":
			valStack.Push(num+valStack.Pop().(float64))
		case "-":
			valStack.Push(valStack.Pop().(float64)-num)
		case "*":
			valStack.Push(num*valStack.Pop().(float64))
		case "/":
			valStack.Push(valStack.Pop().(float64)/num)
		}
		if len(opsStack.el)==0 {
			break
		}
	}
	fmt.Println("操作栈为",valStack.Pop())
}