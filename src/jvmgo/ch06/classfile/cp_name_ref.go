package classfile
    
/*
    xinxin.shi
    2017-06-04 00:48:40
*/
type ConstantMemberrefInfo struct {
    cp                  ConstantPool
    classIndex          uint16
    nameAndTypeIndex    uint16
}       

func (self *ConstantMemberrefInfo) readInfo (reader *ClassReader) {
    self.classIndex = reader.readUint16()
    self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
    return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
    return self.cp.getNameAndType(self.nameAndTypeIndex)
}

// new therr struct extends ConstantMemberrefInfo
type ConstantFieldrefInfo struct {
    ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
    ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
    ConstantMemberrefInfo
}