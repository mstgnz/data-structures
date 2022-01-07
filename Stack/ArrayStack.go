package Stack

import "fmt"

type IArrayStack interface{
	Push(data int)
	Pop()
	Print()
	List() []int
}

type arrayStack struct {
	Arr []int
	ArrSize int
	Index int
}

func ArrayStack() IArrayStack{
	return &arrayStack{[]int{0,0}, 2,0}
}

// Constructor initialize (manuel) -> but interface with auto constructor.
/*func (arr *arrayStack) Constructor(){
	arr.ArrSize = 2
	arr.Index = 0
	arr.Arr = []int{0,0}
}*/

// Push Add to data
func (arr *arrayStack) Push(data int) {
	if arr.Index >= arr.ArrSize{
		newArr := make([]int, arr.ArrSize * 2)
		for i := 0; i < arr.ArrSize; i++ {
			newArr[i] = arr.Arr[i]
		}
		arr.Arr = newArr
		arr.ArrSize *= 2
	}
	arr.Arr[arr.Index] = data
	arr.Index++
}

// Pop Remove to data
func (arr *arrayStack) Pop() {
	arr.Index--
	arr.Arr[arr.Index] = 0
	if arr.Index <= arr.ArrSize / 4{
		newArr := make([]int, arr.ArrSize / 2)
		for i := 0; i < arr.Index; i++ {
			newArr[i] = arr.Arr[i]
		}
		arr.Arr = newArr
		arr.ArrSize /= 2
	}
}

// List - data slice
func (arr *arrayStack) List() []int{
	var list []int
	for i := 0; i < arr.Index; i++ {
		list = append(list, arr.Arr[i])
	}
	return list
}

// Print data
func (arr *arrayStack) Print() {
	fmt.Print("print : ")
	for _, val := range arr.List() {
		fmt.Print(val," ")
	}
	fmt.Println()
}