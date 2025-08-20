package main

import (
    "fmt"
    ArrayList "go-utils/utils/collections"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
    first := ArrayList.NewArrayList[int]()
    first.Add(1)
    first.Add(2)
    first.Add(3)
    fmt.Println("first array:", first.Values())

    second := ArrayList.NewArrayList[int]()
    second.AddAll(first.Values())
    if err := second.InsertAt(2, 13); err != nil {
        panic(err)
    }
    fmt.Println("second array:", second.Values())

    third := ArrayList.NewArrayList[string]()
    third.Add("a")
    third.Add("b")
    third.Add("c")
    fmt.Println("third array:", third.Values())

    third.Clear()
    fmt.Println("third array after clear:", third.Values())
}
