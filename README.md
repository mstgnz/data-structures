# Data Structures With Go

This repository explores various data structures implemented in the Go programming language.


### Linked List

![Linear Linked List](img/LinearLinkedList.png)
```go
type linear struct {
    Data    int
    Next *linear
}
```

![Circular Linked List](img/CircularLinkedList.png)
```go
type circular struct {
    Data    int
    Next *circular
}
```

![Double Linked List](img/DoubleLinkedList.png)
```go
type double struct {
    Data    int
    Next *double
    Prev *double
}
```

## Queue
- Array Queue
```go
type arrayQueue struct {
    Arr []int
    ArrSize int
    FirstIndex int
    LastIndex int
}
```
- Linked List Queue
```go
type linkedListQueue struct {
    X int
    Next *linkedListQueue
}
```

## Stack
- Array Stack
```go
type arrayStack struct {
    Arr []int
    ArrSize int
    Index int
}
```
- Linked List Stack
```go
type linkedListStack struct {
    X int
    Next *linkedListStack
}
```