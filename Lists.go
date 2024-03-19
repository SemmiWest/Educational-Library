package main

import "fmt"

type Block struct {
	data int
	next *Block
}

type LinkedList struct {
	head *Block
}

func (list *LinkedList) InsertInBack(data int) {
	newBlock := &Block{data: data, next: nil}

	if list.head == nil {
		list.head = newBlock
		return
	}
	position := list.head
	for position != nil {
		position = position.next
	}
	position.next = newBlock
}

func (list *LinkedList) InsertInStart(data int) {
	if list.head == nil {
		newBlock := &Block{data, nil}
		list.head = newBlock
		return
	}
	newBlock := &Block{data, list.head}
	list.head = newBlock

}

func (list *LinkedList) InsertAfterBlock(dataLastBlock, newData int) {

	newBlock := &Block{newData, nil}

	after := list.head

	for after != nil {
		if after.data == dataLastBlock {
			newBlock.next = after.next
			after.next = newBlock
			return
		}
		after = after.next
	}

	fmt.Println("Couldn't find the required block")

}

func (list LinkedList) InsertBeforeBlock(newData, varData int) {

	newBlock := &Block{newData, nil}
	point := list.head

	if point == nil {
		return
	}

	if list.head.data == varData {
		newBlock.next = point
		return
	}

	for point != nil {
		if point.next.data == varData {
			newBlock.next = point.next
			point.next = newBlock
			return
		}
		point = point.next
	}
	fmt.Println("Couldn't find the required block")

}

func (list *LinkedList) DeleteBlock(data int) {

	point := list.head

	if point.data == data {
		list.head = point.next
		return
	}
	secondPoint := point.next
	for secondPoint != nil {
		if secondPoint.next.data == data {
			point.next = secondPoint.next
		}
		secondPoint = nil
		return
	}

	fmt.Println("Couldn't find the required block")

}

func (list *LinkedList) CountBlock() (count int) {

	point := list.head
	for point != nil {
		count++
		point = point.next
	}
	return
}

func (list *LinkedList) SearchIndexBlock(index int) *Block {
	count := 0
	point := list.head

	for point != nil {
		count++
		point = point.next
	}
	point = list.head

	if index <= 0 || index > count {
		return nil
	}
	for i := 0; i < count; i++ {
		point = point.next
	}
	return point
}
