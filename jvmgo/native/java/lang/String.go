package lang

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

func init() {
	_string(intern, "intern", "()Ljava/lang/String;")
}

func _string(method func(frame *rtda.Frame), name, desc string) {
	rtc.RegisterNativeMethod("java/lang/String", name, desc, method)
}

// public native String intern();
// ()Ljava/lang/String;
func intern(frame *rtda.Frame) {
	vars := frame.LocalVars()
	jStr := vars.GetThis()

	goStr := rtda.GoString(jStr)
	internedStr := rtda.InternString(goStr, jStr)

	stack := frame.OperandStack()
	stack.PushRef(internedStr)
}
