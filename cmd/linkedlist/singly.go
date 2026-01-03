package main

import "fmt"

// Define structure
type SinglyListNode struct {
	Val  int
	Next *SinglyListNode
}

func main() {

	// Create empty linkedlist
	list := &SinglyListNode{}

	// Create the first node

	list = Prepend(list, 4)

	list = Append(list, 5)

	//Append nodes to list
	list = Append(list, 10)

	list = Append(list, 15)

	list = Append(list, 20)

	showList(list)

	//delete(list, 5)

	showList(list)

	list = reverserList(list)

	showList(list)

	showList(middleNode(list))
}

func middleNode(listHead *SinglyListNode) *SinglyListNode {
	//Using Two Pointers (fast and slow) approach in LinkedList
	//Basicaly, by doing the ahead pointer step by 2, when it got the end the current will be in the middle
	//Start both at the same position
	current := listHead
	ahead := listHead

	//Without both condition, we can have scenarios where ahead.next.next simply does not exit - leading do crashloop
	for ahead != nil && ahead.Next != nil {

		ahead = ahead.Next.Next
		current = current.Next
	}

	return current

}

func reverserList(listHead *SinglyListNode) *SinglyListNode {
	if listHead == nil {
		return listHead
	}

	//got head, aux var help maintain head track
	current := listHead
	//prev aux var
	var prev *SinglyListNode = nil

	for current != nil {
		//save next node address reference in next aux var
		next := current.Next

		// update next node address reference, to previous node address, start null, then is the current node address) in the list
		current.Next = prev
		// update prev var with current node address
		prev = current
		// increase index
		current = next

	}
	// return list head, which is prev, because since we are reversing, it now points to the "end" of the list
	return prev
}

func delete(listHead *SinglyListNode, val int) *SinglyListNode {

	//if the list is empty, return nil
	if listHead == nil {
		return nil
	}

	//if head is to be deleted, return its next node address, which mean that this node will be the new head or the next is nil and the list is now empty
	if listHead.Val == val {
		return listHead.Next
	}

	// use aux to not lose head pointer reference
	current := listHead

	// jump head, from current.Next looks if the next node exist
	for current.Next != nil {
		// if so
		// check the val of the node with address defined in current.Next is the deletion target
		if current.Next.Val == val {
			//To delete, change the address of the next node on the current node, jumping the next node since it is deletion target
			// It still in memory... How Go handles it?
			current.Next = current.Next.Next
			//return head pointer
			return listHead
		}
		//increase loop index
		current = current.Next
	}

	return listHead
}

func Prepend(listHead *SinglyListNode, val int) *SinglyListNode {

	// Create node with the desired val, and next as head
	newNode := &SinglyListNode{Val: val, Next: listHead}

	if listHead == nil {
		return newNode
	}

	current := listHead
	newNode.Next = current

	return newNode
}

func Append(listHead *SinglyListNode, val int) *SinglyListNode {

	//Create node struct, store address to var
	newNode := &SinglyListNode{Val: val, Next: nil}

	//Check if list is empty, if so, return newNode as the linkedlist
	if listHead == nil {
		return newNode
	}

	// Use auxiliar variable to not lose head refrence
	current := listHead

	// iterate linkedlist and got the last node
	for current.Next != nil {
		current = current.Next

	}
	// with the pointer to the last node, access it in .Next and pass the newNode pointer address
	current.Next = newNode
	return listHead

}

func showList(list *SinglyListNode) {
	// use auxiliar variable to iterate linkedlist
	current := list

	for current != nil {
		fmt.Printf("%d ->", current.Val)
		current = current.Next

	}
	println()
}
