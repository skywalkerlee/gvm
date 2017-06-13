package constants

import "github.com/skywalkerlee/gvm/instructions/base"
import "github.com/skywalkerlee/gvm/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
