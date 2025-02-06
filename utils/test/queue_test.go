package utils

import (
	"testing"

	"example.com/ogrex/utils"
)

func TestShouldCreateQueueWithGibenInitialElems(t *testing.T) {
	queue := utils.NewFifoQueue(1, 2, 3)
	wantedLength := 3
	receivedLength := queue.GetLength()

	if receivedLength != wantedLength {
		t.Fatalf("Wrong queue length. Length given: %d Length wanted: %d", receivedLength, wantedLength)
	}
}

func TestShouldTakeLastElementFromQueue(t *testing.T) {
	queue := utils.NewFifoQueue(1, 2, 3)
	wantedLength := 2
	wantedElem := 3

	elem, _ := queue.TakeLast()

	if wantedLength != queue.GetLength() {
		t.Fatalf("Wrong length after take. Wanted length: %d Received length: %d", wantedLength, queue.GetLength())
	}

	if elem != wantedElem {
		t.Fatalf("Wrong elem taken. Wanted elem: %d Received elem: %d", wantedElem, elem)
	}
}

func TestShouldInsertFirstElementIntoQueue(t *testing.T) {
	queue := utils.NewFifoQueue(1)
	wantedLength := 2
	wantedElem := 2

	queue.Insert(2)

	if queue.GetLength() != 2 {
		t.Fatalf("Wrong length after insert. Wanted length: %d Received length: %d", wantedLength, queue.GetLength())
	}

	firstElem, _ := queue.TakeLast()
	_ = firstElem

	secondElem, _ := queue.TakeLast()

	if secondElem != wantedElem {
		t.Fatalf("Queue returning elems in wrong order. Wanted elem: %d Received elem: %d", wantedElem, secondElem)
	}
}

func TestShouldPeekLastElemWithoutRemovingIt(t *testing.T) {
	queue := utils.NewFifoQueue(1, 2, 3)
	queue.Insert(4)
	wantedElem := 3
	wantedLength := 4

	elem, _ := queue.Peek()

	if wantedLength != queue.GetLength() {
		t.Fatalf("Length shouldnt be reduced after .Peek() method. Wanted length: %d Received length: %d", wantedLength, queue.GetLength())
	}

	if elem != wantedElem {
		t.Fatalf("Wrong elem peeked. Wanted elem: %d Received elem: %d", wantedElem, elem)
	}
}

func TestShouldThrowErrorIfQueueIsEmpty(t *testing.T) {
	queue := utils.NewFifoQueue[int]()
	_, err := queue.TakeLast()

	if err == nil {
		t.Fatal("Error not returned when queue was empty.")
	}
}
