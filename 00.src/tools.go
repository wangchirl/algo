package src

import (
	"fmt"
	"testing"
)

//断言函数, 简单快捷好用
func So(t *testing.T, expect, actual interface{}) {
	if expect != actual {
		t.Errorf("%s, expect:%v, actual:%v\n", t.Name(), expect, actual)
	}
}
//断言函数,带name
func Asset(name interface{}, t *testing.T, expect, actual interface{}) {
	var nameStr string
	switch name.(type) {
	case int:
		nameStr = fmt.Sprint(name)
	case string:
		nameStr = fmt.Sprint(name)
	default:
		nameStr = "default"
	}
	t.Run(nameStr, func(t *testing.T) {
		if expect != actual {
			t.Errorf("%s, expect:%v, actual:%v\n", t.Name(), expect, actual)
		}
	})
}

type ToolsInteger struct {
}
func NewInt() *ToolsInteger {
	return &ToolsInteger{}
}
//切片加入元素, index是插入位置, value是需要插入的值
func (ToolsInteger) InsertIndex(arr []int, index, value int) []int {
	rear := append([]int{}, arr[index:]...)
	return append(append(arr[:index], value), rear...)
}
//切片克隆
func (ToolsInteger) Clone(src []int) (dst []int) {
	dst = make([]int, len(src))
	copy(dst,src)
	return
}
type ToolsByte struct {
}
func NewByte() *ToolsByte {
	return &ToolsByte{}
}
//切片加入元素, index是插入位置, value是需要插入的值
func (ToolsByte) InsertIndex(arr []byte, index, value byte) []byte {
	rear := append([]byte{}, arr[index:]...)
	return append(append(arr[:index], value), rear...)
}
//切片克隆
func (ToolsByte) Clone(src []byte) (dst []byte) {
	dst = make([]byte, len(src))
	copy(dst,src)
	return
}
