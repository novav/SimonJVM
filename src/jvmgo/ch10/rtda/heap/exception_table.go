package heap

/*
    Chap10.3
    异常处理表
    xinxin.shi
    2017-09-05 23:25:06
*/
import "jvmgo/ch10/classfile"

type ExceptionTable []*ExceptionTable

type  ExceptionTable struct {
    startPc    int
    endPc      int
    handlerPc  int
    catchType  *ClassRef
}

func newExceptionTable(entries []*classfile.ExceptionTableEntry,
    cp *ConstantPool) ExceptionTable {
    table := make([]*ExceptionHandler, len(entries))
    for i, entry := range entries {
        table[i] = &ExceptionHandler{
            startPc:      int(entry.StartPc()),
            endPc:        int(entry.EndPc()),
            handlerPc:    int(entry.HandlerPc()),
            catchType:    getCatchType(uint(entry.CatchType()), cp),
        }
    }
}

func getCatchType(index uint, cp *ConstantPool) *ClassRef {
    if index == 0 {
        return nil
    }
    return cp.GetConstant(index).(*ClassRef)
}

func (self *ExceptionTable)findExceptionHandler(exClass *Class,
    pc int) *ExceptionHandler {
    for _, handler := range self {
        if pc >= handler.startPc && pc < handler.endPc {
            if handler.catchType == nil {
                return handler // catch all
            }
            catchClass := handler.catchType.ResolvedClass()
            if catchClass == exClass || catchClass.IsSuperClassOf(exClass) {
                return handler
            }
        }
    }
    return nil
}