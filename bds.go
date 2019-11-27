package BDS

// basic data structure interface include
// two method named Pop and Push
type BDS interface {
	Pop() interface{}
	Push(data interface{})
}
