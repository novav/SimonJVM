package heap 
import "fmt"
import "jvmgo/ch06/classfile"
/*
    Chap 6.1
    运行时常量池
    xinxin.shi
    2017-07-27 23:55:44
*/
type Constant interface{}
type ConstantPool struct {
    class  *Class 
    consts []Constant 
}

/*核心逻辑就是把[]classfile.ConstantInfo转换成[]heap.Constant*/
func newConstantPool(class *Class, cfCp classfile.ConstantPool ) *ConstantPool {
    cpCount := len(cfCp)
    consts := make([]Constant, cpCount)
    rtCp := &ConstantPool{class, consts}
    for i := 1; i < cpCount; i++ {
        cpInfo := cfCp[i]
        switch cpInfo.(type) {
        case *classfile.ConstantIntegerInfo:
            intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
            consts[i] = intInfo.Value() // int32
        case *classfile.ConstantFloatInfo:
            floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
            consts[i] = floatInfo.Value() // float32
        case *classfile.ConstantLongInfo:
            longInfo := cpInfo.(*classfile.ConstantLongInfo)
            consts[i] = longInfo.Value() // int64
            i++
        case *classfile.ConstantDoubleInfo:
            doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
            consts[i] = doubleInfo.Value() // float
            i++
        case *classfile.ConstantStringInfo:
            stringInfo := cpInfo.(*classfile.ConstantStringInfo)
            consts[i] = stringInfo.Value() // string

        case *classfile.ConstantClassInfo:
            classInfo := cpInfo.(*classInfo.ConstantClassInfo)
            consts[i] = newClassRef(rtCp, classInfo) // 6.2.1
        case *classfile.ConstantFieldrefInfo:
            fieldrefInfo := cpInfo.(*classfile.ConstantFieldrefInfo)
            consts[i] = newFieldRef(rtCp, fieldrefInfo) // 6.2.2
        case *classFile.ConstantMethodrefInfo:
            methodrefInfo := cpInfo.(*classfile.ConstantMethodrefInfo)
            consts[i] = newMethodRef(rtCp, methodrefInfo) // 6.2.3
        case *classfile.ConstantInterfaceMethodrefInfo:
            methodrefInfo := cpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
            consts[i] = newInterfaceMethodRef(rtCp, methodrefInfo) // 6.2.4
        }
    }
    return rtCp    
}

func (self *ConstantPool) GetConstant(index uint) Constant {
    if c := self.consts[index]; c != nil {
        return c
    }
    panic(fmt.Sprintf("No constant at index %d", index))
}
