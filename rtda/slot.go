package rtda

import "github.com/skywalkerlee/gvm/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
