package stack

import "github.com/skywalkerlee/gvm/instructions/base"
import "github.com/skywalkerlee/gvm/rtda"

// Swap the top two operand stack values
type SWAP struct{ base.NoOperandsInstruction }

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
