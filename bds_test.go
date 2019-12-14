package BDS

import "testing"

import "fmt"

func TestArray(t *testing.T) {
	a := NewArray(1)
	d, err := a.Push()
	a.Pop(12)
	d, err = a.Push()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)
}
