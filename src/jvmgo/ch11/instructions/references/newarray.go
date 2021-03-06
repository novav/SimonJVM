package references

import "jvmgo/ch11/instructions/base"
import "jvmgo/ch11/rtda"
import "jvmgo/ch11/rtda/heap"

/*
    Chap 8.3.1
    anewarray指令
    xinxin.shi
    2017-08-17 23:39:25
*/
const (
    AT_BOOLEAN      = 4
    AT_CHAR         = 5
    AT_FLOAT        = 6
    AT_DOUBLE       = 7
    AT_BYTE         = 8
    AT_SHORT        = 9
    AT_INT          = 10
    AT_LONG         = 11
)

// Create new array of proimitive

type NEW_ARRAY struct {
    atype uint8
}


func (self *NEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
    self.atype = reader.ReadUint8()
}

func (self *NEW_ARRAY) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    count := stack.PopInt()
    if count < 0 {
        panic("java.lang.NegativeArraySizeException")
    }
    classLoader := frame.Method().Class().Loader()
    arrClass := getPrimitiveArrayClass(classLoader, self.atype)
    arr := arrClass.NewArray(uint(count))
    stack.PushRef(arr)
}


func getPrimitiveArrayClass(loader *heap.ClassLoader, atype uint8) *heap.Class{
    switch atype {
        case AT_BOOLEAN:    return loader.LoadClass("[Z")
        case AT_BYTE:    return loader.LoadClass("[B")
        case AT_CHAR:    return loader.LoadClass("[C")
        case AT_SHORT:    return loader.LoadClass("[S")
        case AT_INT:    return loader.LoadClass("[I")
        case AT_LONG:    return loader.LoadClass("[J")
        case AT_FLOAT:    return loader.LoadClass("[F")
        case AT_DOUBLE:    return loader.LoadClass("[D")
        default: panic("Invalid atype!")
    }
}