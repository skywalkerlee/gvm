package control

import "github.com/skywalkerlee/gvm/instructions/base"
import "github.com/skywalkerlee/gvm/rtda"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
