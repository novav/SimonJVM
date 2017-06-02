package classfile

/*  xinxin.shi
    2017-06-02
*/

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
    cpCount := int(reader.readerUint16())
    cp := make([]ConstantInfo, cpCount)

    for i := 1; i < cpCount; i++ { // index start from 1
        cp[i] = readConstantInfo(reader, cp)
        switch cp[i].(type) {
        case *ConstantLongInfo, *ConstantDoubleInfo:
            i++ // 站两个位置
        }
    }
    return cp
}

func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
    if cpInfo := self[index]; cpInfo != nil {
        return cpInfo
    }
    panic("Invalid constant pool index!")
}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {
    ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
    name := self.getUtf8(ntInfo.nameIndex)
    _type := self.getUtf8(ntInfo.descriptorIndex)
}

func (self ConstantPool) getClassName(index uint16) string {
    classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
    return self.getUtf8(classInfo.nameIndex)
}

func (self ConstantPool) getUtf8(index uint16) string {
    utfInfo := self.getConstantInfo(index).(*ConstantUtf8Info)
    return utf8Info.str
}