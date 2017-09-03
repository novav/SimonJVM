package lang

import "math"
import "jvmgo/ch09/native"
import "jvmgo/ch09/rtda"

/*
    Chap9.4.3
    Float.floatToRawIntBits
    xinxin.shi
    2017-08-30 21:31:59
*/
func init() {
    native.Register("java/lang/Float", "floatToRawIntBits", "(F)I", floatToRawIntBits)
}

// public static native int floatToRawIntBits(float value);
func floatToRawIntBits(frame *rtda.Frame) {
    value := frame.LocalVars().GetFloat(0)
    bits := math.Float32bits(value)
    frame.OperandStack().PushInt(int32(bits))
}