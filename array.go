package BDS

import "fmt"

type Array struct {
	Value       []interface{} `json:"value"`
	LengthValue int           `json:"length"`
}

// creat Array struct.
func NewArray(num int) *Array {
	result := make([]interface{}, num)
	return &Array{
		Value:       make([]interface{}, num),
		LengthValue: len(result),
	}
}

// pop a data.
func (a *Array) Pop(data interface{}) {
	a.Value = append(a.Value, data)

}

// pubsh a data.
func (a *Array) Push() (d interface{}, err error) {

	if a.Length() <= 0 {
		return nil, fmt.Errorf("error:cant push data,length is 0")
	} else {
		d = a.Value[len(a.Value)-1]
		a.Value = a.Value[:a.Length()-1]
	}

	return
}

// return array's real length by len(value).
func (a *Array) Length() int {
	return len(a.Value)
}
