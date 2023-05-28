package main

import (
	"container/list"
	"fmt"
)


// doubly linked list using the "list" standard package


// A doubly linked list is a data structure that consists of a sequence of elements, 
//where each element contains a value and maintains references to the previous and next elements in the list.

func main() {
	var exList list.List  // create a variable of type list.List to hold the the doubly linked list


	// data is added to the list using .PushBack()
	exList.PushBack(23)
	exList.PushBack(45)
	exList.PushBack(98)


	//looping through the list
	for element := exList.Front();element != nil;element = element.Next(){
		fmt.Println(element.Value.(int)) // print the value of the current element
	}
}