package lang

import "math"
import "jvmgo/ch11/native"
import "jvmgo/ch11/rtda"

const jlFloat = "java/lang/Float"
/*
    Chap9.4.3
    Float.floatToRawIntBits
    xinxin.shi
    2017-08-30 21:31:59
*/
func init() {
	native.Register(jlFloat, "floatToRawIntBits", "(F)I", floatToRawIntBits)
	native.Register(jlFloat, "intBitsToFloat", "(I)F", intBitsToFloat)
}

// public static native int floatToRawIntBits(float value);
func floatToRawIntBits(frame *rtda.Frame) {
    value := frame.LocalVars().GetFloat(0)
    bits := math.Float32bits(value)
    frame.OperandStack().PushInt(int32(bits))
}

// public static native float intBitsToFloat(int bits);
// (I)F
func intBitsToFloat(frame *rtda.Frame) {
	bits := frame.LocalVars().GetInt(0)
	value := math.Float32frombits(uint32(bits)) // todo
	frame.OperandStack().PushFloat(value)
}
