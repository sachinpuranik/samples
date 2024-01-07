package bar

import (
	"fmt"
)

// MyType -
type MyType struct{}

// MyPrint -
func (m MyType) MyPrint() {
	fmt.Println("Wow Sachin")
}

// New -
func New() *MyType {
	return &MyType{}
}
