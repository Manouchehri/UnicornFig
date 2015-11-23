package stdlib

import (
	uni "../interpreter"
	//"errors"
)

func SLIB_List(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	list := uni.NewList()
	if len(arguments) == 0 {
		return nil, list, env
	}
	lastNode := list.List
	for _, value := range arguments {
		lastNode.Next = &uni.ListNode{}
		lastNode.Next.Data = value
		lastNode.Next.Next = nil
		lastNode = lastNode.Next
	}
	return nil, list, env
}
