package references

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

// Get static field from class
type GET_STATIC struct {
	base.Index16Instruction
	field *rtc.Field
}

func (self *GET_STATIC) Execute(frame *rtda.Frame) {
	if self.field == nil {
		cp := frame.Method().Class().ConstantPool()
		kFieldRef := cp.GetConstant(self.Index).(*rtc.ConstantFieldref)
		self.field = kFieldRef.StaticField()
	}

	class := self.field.Class()
	if class.InitializationNotStarted() {
		frame.RevertNextPC() // undo getstatic
		frame.Thread().InitClass(class)
		return
	}

	val := self.field.GetStaticValue()
	stack := frame.OperandStack()
	stack.PushField(val, self.field.IsLongOrDouble)
}
