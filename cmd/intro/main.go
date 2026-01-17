package main

import (
	"fmt"
)

func main() {
	// THe blocks to build a data structure are very similar, and mainly the difference is the way that we add, remove and find elements.
	// Beside that we are always using arrays and linked lists.

	graphs()
	heap()
	binaryTree()
	stack()
	hashMap()
	queue()
	linkedList()
	array()
	bigo()

}

func trie() {

	/*
		Trie - Data Structure // A type of tree (also known as prefix tree)
		- A trie is a tree-like data structure that stores a dynamic set of strings.
		- Each node in a trie represents a single character.
		- The root node represents an empty string.
		- It retains the prefix property, meaning that each node's children share a common prefix of the string associated with the previous node.

		Tries are used to store dictionaries, phone directories, and autocomplete features.

	*/
}

func graphs() {
	fmt.Println("Graphs")
	/*
		Graphs
		- A graph is a collection of nodes(vertices) and edges that connect the nodes.
		- A graph can be directed or undirected.
		- In a directed graph, the edges have a direction.
		- In an undirected graph, the edges do not have a direction.
		- A graph can be weighted or unweighted.
		- In a weighted graph, each edge has a weight.
		- In an unweighted graph, each edge does not have a weight.
	*/
}

func heap() {
	fmt.Println("Heap")

	/*
		Heap - A special tree-based data structure that satisfies the heap property.
		- A heap is a complete binary tree.
		- A heap can be a min-heap or a max-heap.
		- In a min-heap, the parent node is less than or equal to the child nodes.
		- In a max-heap, the parent node is greater than or equal to the child nodes.

	*/

}
func binaryTree() {
	fmt.Println("Binary Tree")

	/*
		Binary Tree
		- A tree data structure where each node has at most two children.
		- The top node is called the root node(head).
		- The nodes at the bottom are called leaf nodes.
		- The nodes that are not leaf nodes are called internal nodes.
		- The height of the tree is the number of edges on the longest path from the root node to a leaf node.
		- The depth of a node is the number of edges on the path from the root node to the node.
		- The height of the tree is the height of the root node.
		- The depth of the root node is 0.
		- The height of an empty tree is -1.

		Binary Search Tree (BST)
		- A binary tree where the left child is less than the parent node and the right child is greater than the parent node.
		- The left subtree of a node contains only nodes with keys less than the node's key.
		- The right subtree of a node contains only nodes with keys greater than the node's key.
		- The left and right subtree must also be a binary search tree.

		Balanced Binary Tree (B-Tree)
		- A binary tree where the height of the left and right subtrees of any node differ by at most one.
		- The height of the left and right subtrees of the root node differ by at most one.
		- The left and right subtrees are also balanced binary trees.

		Complete Binary Tree
		- A binary tree where all levels are completely filled except possibly for the last level, which is filled from left to right.
		- A complete binary tree is a balanced binary tree.

		Full Binary Tree
		- A binary tree where each node has either zero or two children.
		- A full binary tree is a complete binary tree.

		Perfect Binary Tree
		- A binary tree where all internal nodes have two children and all leaf nodes are at the same level.
		- A perfect binary tree is a full binary tree.

		Traversal
		- Pre-order Traversal: Visit the root node first, then the left subtree, and finally the right subtree.
		- In-order Traversal: Visit the left subtree first, then the root node, and finally the right subtree.
		- Post-order Traversal: Visit the left subtree first, then the right subtree, and finally the root node.
		- Level-order Traversal: Visit nodes level by level from left to right.

	*/

}
func stack() { //also know as pilha in portuguese
	fmt.Println("Stack")

	/*
		LIFO - Last In First Out
		- Implemented using arrays or linked lists.
		- The last element added to the stack will be the first to be removed.
		- An pointer to the top element is maintained to add and remove elements in O(1) time complexity.
	*/
}

func hashMap() {
	fmt.Println("Hash Map")
	/*

		An hash is made on key to generate an index in an array where the value will be stored,
		sometime this values are structs or objects, usually a small linkedlist that contains
		different elements with the same index. By doing that we keep time complexity O(1) while handle collisions.

		Hash Map - Key-Value Pair Data Structure (basicaly a dictionary)
		Usually the complexity to add, remove and find an element is O(1) - constant time.

		The hash function is used to convert the key into an index in the array, since it is in a contiguous memory location, we can access the element in O(1) time complexity.
		The hash functionon probably will be a SHA-256 or MD5, but it can be any function that converts the key into an index in the array.
		A modulo operation is used to ensure that the index (SHA or MD5) is within the array bounds.


		Collision can happen when two keys are converted into the same index, so we need to handle it.
		There are two ways to handle collisions:
		- Open Addressing: When a collision happens, we find the next available index in the array. It is not the best approach because in the worst case scenario we would need to iterate over all elements to find the next available index.
		- Separate Chaining: When a collision happens, we store the elements in a linked list. It is the best approach because we can have multiple elements in the same index, this index will have a linked list with all elements that have the same index.
		Since this sublist its supposed to be small, the time complexity to find an element is O(1).

		Load Factor = Number of elements / Number of buckets. It defines how full the hash map is and it is used to determine when to resize the hash map.
		Keep the load factor below 0.7 to avoid collisions. When the load factor is above 0.7, we need to resize (double) the hash map to avoid collisions.


	*/
}
func queue() {

	fmt.Println("Queue")

	/*
		FIFO - First In First Out
		- The first element added to the queue will be the first to be removed.
		- The traditional implementation of a queue is done using linked lists, however, we can also use arrays.


		Queues are data structres utilized for buffering, task scheduling, data streaming, etc.

		If correct implemented, the queue will maintain an reference to the first(head) and last(tail) element, so we can add and remove elements in O(1) time complexity.
		If we don't have the reference to the last element, we would need to iterate over all elements to find the last element, which would be O(n) time complexity.

		In a dequeue (double linked list), we can add and remove elements from the head and tail - it is not necessarily a FIFO structure.
	*/
}

func linkedList() {
	fmt.Println("Linked List")
	/*
			It is not necessary to define the size of the linked list, because each element has a pointer to the next element.
			Which also mean that the elements are not allocated in contiguous memory locations.
			Types of linked lists:
			- Singly Linked List: Each element has a pointer to the next element.
			- Doubly Linked List: Each element has a pointer to the next and previous element.
			- Circular Linked List: The last element points to the first element and the first element points to the last element.

			Linked List vs Array:
			- Read: Array is faster than Linked List because the elements are allocated in contiguous memory locations. The compplexity to found an element in an array is O(1) and in a linked list is usualy O(n).
			- Insert: Linked List is faster than Array because we don't need to move all elements (basicaly recreate the array) to insert a new element.
			- Delete: Linked List is faster than Array because we don't need to move all elements (basicaly recreate the array) to delete an element.

			Since in array it is in contiguous memory locations, we can calculate the memory address of the element by the index just by using the formula: memory_address = base_address + (index * size_of_element).
			In the linked list, to find it lenght, for example, we need to iterate over all the elements, the same happen to find an specific element.

		 	The scenarios will depend if the changes are in the head/tail or in the middle of the list
			#	Read: O(1) in the best case and O(n) in the worst case
			#	Insert: O(1) in the best case and O(n) in the worst case
			#	Delete: O(1) in the best case and O(n) in the worst case

	*/
}

func array() {
	fmt.Println("Array")
	arr := [5]uint16{1, 2, 3, 4, 5}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Element %d: %d, Address: %p\n", i, arr[i], &arr[i])
	}

	/*
	   Array - Fixed-size collection of elements of the same type allocated in contiguous memory locations.
	   Its important define the right size of the array, because it cannot be changed later.
	   It cannot be changed because when defining the type (int, string, in32, etc) and size the memory is allocated in a contiguous way.
	   Which means that maybe we have other data in the memory after/before the array, so we cannot change the size of the array because we would be overwriting other data in the memory.
	*/

}

func bigo() {
	fmt.Println("Big O Notation")
	/*
	   Big O Notation - Measuring Algorithm scale and Performance

	   - Big O notation is a way to describe the scale of an algorithm, but it does not necessarily represent execution time (performance in terms of speed).
	   - It focuses on how the execution time (or space usage) grows as the input size increases.

	   - Without asymptotic analysis, it’s impossible to determine if an O(n) algorithm is always faster than an O(n²) algorithm.
	     However, we can predict which one will scale better as input size increases.

	   - Big O notation only describes how well an algorithm scales with input size, not how fast it runs in absolute terms.

	   - It is used to measure:
	     - **Time Complexity** (Execution Time)
	     - **Space Complexity** (Memory Usage)

	   O(1) - Constant Time & Constant Space
	   Example: Accessing an array element by index or retrieving the first element of a list.
	   - No matter how large the input is, the execution time remains constant.
	   - Worst-case scenario: execution time does not change with input size.

	   O(log N) - Logarithmic Time
	   Example: Binary Search
	   - As the input size increases exponentially, the execution time increases only linearly.
	   - Each step reduces the problem size significantly (typically by half).

	   O(N) - Linear Time & Linear Space
	   Example: Looping through an array
	   - Execution time grows proportionally with input size.

	   O(N log N) - Linearithmic Time
	   Example: Quick Sort, Merge Sort
	   - The array is divided recursively (O(log N)), and each level of recursion processes all elements (O(N)).
	   - This results in an overall complexity of O(N log N).

	   O(N²) - Quadratic Time
	   Example: Bubble Sort, Selection Sort (Nested loops)
	   - For each element, the algorithm iterates over the entire array again.
	   - Execution time grows quadratically as input size increases.

	   O(2^N) - Exponential Time
	   Example: Recursive Fibonacci Algorithm
	   - Every function call branches into two more calls, leading to an exponential number of recursive calls.
	   - Example: Fibonacci using naive recursion (`fib(n) = fib(n-1) + fib(n-2)`).

	   O(N!) - Factorial Time
	   Example: Traveling Salesman Problem (brute force solution)
	   - Worst-case scenario: The algorithm tries **every possible** ordering of N elements.
	   - Example: Finding the shortest route visiting N cities by checking all (N!) permutations.

	*/
}
