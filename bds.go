package BDS

var (
	typeName map[string]BDS
)

// basic data structure interface include
// two method named Pop and Push
type BDS interface {
	Pop(data interface{})                 // 输入数据
	Push() (value interface{}, err error) // 输出数据
	Length() int                          // 输出它的长度
}
