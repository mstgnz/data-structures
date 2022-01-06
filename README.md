# Data Structures With Go

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

### Source
[![](https://www.sadievrenseker.com/logo1_copy.png)](https://bilgisayarkavramlari.com/category/veri-yapilari/)
<a href="https://www.geeksforgeeks.org/data-structures/linked-list/"><img src="https://media-exp1.licdn.com/dms/image/C4D22AQHzOgJvJye7CQ/feedshare-shrink_2048_1536/0/1623937238005?e=1643241600&v=beta&t=tdL6PoXMlOlMih6JroVruTdRjtXjsM77FTZnahNnmfo" height="155" />
</a>
