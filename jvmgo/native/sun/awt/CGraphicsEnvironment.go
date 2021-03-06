package awt

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

func init() {
	_cge(cge_initCocoa, "initCocoa", "()V")
	_cge(cge_getMainDisplayID, "getMainDisplayID", "()I")
}

func _cge(method func(frame *rtda.Frame), name, desc string) {
	rtc.RegisterNativeMethod("sun/awt/CGraphicsEnvironment", name, desc, method)
}

func cge_initCocoa(frame *rtda.Frame) {
	//TODO
}

func cge_getMainDisplayID(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}
