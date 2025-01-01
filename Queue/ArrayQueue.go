package Queue

import "fmt"

type IArrayQueue interface {
	Enqueue(data int)
	Dequeue()
	List() []int
	Print()
	reSort()
}

type arrayQueue struct {
	Arr        []int
	ArrSize    int
	FirstIndex int
	LastIndex  int
}

func ArrayQueue() IArrayQueue {
	return &arrayQueue{[]int{0, 0}, 2, 0, 0}
}

// Enqueue Add to data
func (arr *arrayQueue) Enqueue(data int) {
	// Removing an element from the array is just row shifting if before enlarging the array.
	// If the first index is different from 0, the element has been removed from the array
	// Then we will reorder the array instead of enlarging the array.
	// If first index is 0 and last index is bigger than array size we will increase array size
	if arr.LastIndex >= arr.ArrSize {
		if arr.FirstIndex == 0 {
			newArr := make([]int, arr.ArrSize*2)
			for i := 0; i < arr.ArrSize; i++ {
				newArr[i] = arr.Arr[i]
			}
			arr.Arr = newArr
			arr.ArrSize *= 2
		} else {
			arr.reSort()
		}
	}
	arr.Arr[arr.LastIndex] = data
	arr.LastIndex++
}

// Dequeue Remove to data
func (arr *arrayQueue) Dequeue() {
	// if deque is run first
	if arr.FirstIndex == 0 && arr.LastIndex == 0 {
		return
	}
	arr.Arr[arr.FirstIndex] = 0
	arr.FirstIndex++
	if arr.LastIndex-arr.FirstIndex <= arr.ArrSize/4 {
		newArr := make([]int, arr.ArrSize/2)
		sort := 0
		for i := arr.FirstIndex; i < arr.LastIndex; i++ {
			newArr[sort] = arr.Arr[i]
			sort++
		}
		arr.Arr = newArr
		arr.ArrSize /= 2
		arr.LastIndex = arr.LastIndex - arr.FirstIndex
		arr.FirstIndex = 0
	}
}

// ReSort data
func (arr *arrayQueue) reSort() {
	newArr := make([]int, arr.ArrSize)
	sort := 0
	for i := arr.FirstIndex; i < arr.LastIndex; i++ {
		newArr[sort] = arr.Arr[i]
		sort++
	}
	arr.Arr = newArr
	arr.LastIndex = arr.LastIndex - arr.FirstIndex
	arr.FirstIndex = 0
}

// List data - slice
func (arr *arrayQueue) List() []int {
	if arr.FirstIndex >= arr.LastIndex {
		return []int{}
	}
	var list []int
	for i := arr.FirstIndex; i < arr.LastIndex; i++ {
		list = append(list, arr.Arr[i])
	}
	return list
}

// Print data
func (arr *arrayQueue) Print() {
	fmt.Print("print : ")
	for _, val := range arr.List() {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
