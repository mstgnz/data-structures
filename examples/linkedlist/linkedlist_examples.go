package main

import (
	"fmt"
	"strings"

	"github.com/mstgnz/data-structures/linkedlist"
)

func main() {
	// Linear Linked List örnekleri
	fmt.Println("Linear Linked List Örnekleri:")

	// Integer List
	intLinear := linkedlist.NewLinear[int](10)
	intLinear.AddToEnd(20)
	intLinear.AddToEnd(30)
	intLess := func(a, b int) bool { return a < b }
	intLinear.AddToSequentially(15, intLess)
	fmt.Println("\nInteger Linear List:")
	intLinear.Print()

	// String List
	strLinear := linkedlist.NewLinear[string]("Merhaba")
	strLinear.AddToEnd("Dünya")
	strLess := func(a, b string) bool { return strings.Compare(a, b) < 0 }
	strLinear.AddToSequentially("Go", strLess)
	fmt.Println("\nString Linear List:")
	strLinear.Print()

	// Double Linked List örnekleri
	fmt.Println("\nDouble Linked List Örnekleri:")

	// Integer List
	intDouble := linkedlist.NewDouble[int](10)
	intDouble.AddToEnd(20)
	intDouble.AddToEnd(30)
	intDouble.AddToSequentially(15, intLess)
	fmt.Println("\nInteger Double List (İleri):")
	intDouble.Print(false)
	fmt.Println("Integer Double List (Geri):")
	intDouble.Print(true)

	// String List
	strDouble := linkedlist.NewDouble[string]("Merhaba")
	strDouble.AddToEnd("Dünya")
	strDouble.AddToSequentially("Go", strLess)
	fmt.Println("\nString Double List (İleri):")
	strDouble.Print(false)
	fmt.Println("String Double List (Geri):")
	strDouble.Print(true)

	// Circular Linked List örnekleri
	fmt.Println("\nCircular Linked List Örnekleri:")

	// Integer List
	intCircular := linkedlist.NewCircular[int](10)
	intCircular.AddToEnd(20)
	intCircular.AddToEnd(30)
	intCircular.AddToSequentially(15, intLess)
	fmt.Println("\nInteger Circular List:")
	intCircular.Print()

	// String List
	strCircular := linkedlist.NewCircular[string]("Merhaba")
	strCircular.AddToEnd("Dünya")
	strCircular.AddToSequentially("Go", strLess)
	fmt.Println("\nString Circular List:")
	strCircular.Print()

	// Custom struct örneği
	type Person struct {
		Name string
		Age  int
	}

	// Person karşılaştırma fonksiyonları
	personLess := func(a, b Person) bool { return a.Age < b.Age }

	// Linear List with Person
	personLinear := linkedlist.NewLinear[Person](Person{Name: "Ali", Age: 25})
	personLinear.AddToEnd(Person{Name: "Ayşe", Age: 30})
	personLinear.AddToSequentially(Person{Name: "Mehmet", Age: 28}, personLess)
	fmt.Println("\nPerson Linear List:")
	personLinear.Print()

	// Double List with Person
	personDouble := linkedlist.NewDouble[Person](Person{Name: "Ali", Age: 25})
	personDouble.AddToEnd(Person{Name: "Ayşe", Age: 30})
	personDouble.AddToSequentially(Person{Name: "Mehmet", Age: 28}, personLess)
	fmt.Println("\nPerson Double List:")
	personDouble.Print(false)

	// Circular List with Person
	personCircular := linkedlist.NewCircular[Person](Person{Name: "Ali", Age: 25})
	personCircular.AddToEnd(Person{Name: "Ayşe", Age: 30})
	personCircular.AddToSequentially(Person{Name: "Mehmet", Age: 28}, personLess)
	fmt.Println("\nPerson Circular List:")
	personCircular.Print()
}
