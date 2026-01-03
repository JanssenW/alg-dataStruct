package main

import "fmt"

type DoublyListNode struct {
	Val  int
	Next *DoublyListNode
	Prev *DoublyListNode
}

func main() {

	list := &DoublyListNode{}
	list = Append(5, list)
	list = Append(10, list)
	list = Append(15, list)
	showList(list)
	list = Prepend(3, list)
	showList(list)

}

func Prepend(val int, list *DoublyListNode) *DoublyListNode {
	newNode := &DoublyListNode{Val: val, Next: list, Prev: nil}

	if list == nil {
		return newNode
	}

	current := list
	newNode.Next = current
	current.Prev = newNode

	return newNode
}

func Append(val int, list *DoublyListNode) *DoublyListNode {

	newNode := &DoublyListNode{Val: val, Next: nil, Prev: nil}

	if list == nil {
		return newNode
	}

	current := list
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
	newNode.Prev = current

	return list
}

func showList(list *DoublyListNode) {

	current := list
	for current != nil {
		fmt.Printf("%d | ", current.Prev)
		fmt.Printf(" -- %d -- ", current.Val)
		fmt.Printf("%d | ", current.Next)
		fmt.Println("-->>")
		current = current.Next
	}

	println()
}
