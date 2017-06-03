package classfile

/*
cp_info {
    u1 tag;
    u1 info[];
}
*/
/*  xinxin.shi
    2017-06-03 0:38
*/
const {
    CONSTANT_Class                  = 7
    CONSTANT_Fieldref               = 9
    CONSTANT_Methodref              = 10
    CONSTANT_InterfaceMethodref     = 11
    CONSTANT_String                 = 8
    CONSTANT_Integer                = 3
    CONSTANT_Float                  = 4
    CONSTANT_Long                   = 5
    CONSTANT_Double                 = 6
    CONSTANT_NameAndType            = 12
    CONSTANT_Utf8                   = 1
    CONSTANT_MethodHandle           = 15
    CONSTANT_MethodType             = 16
    CONSTANT_InvokeDynamic          = 18
}

type ConstantInfo interface {
    readInfo(reader *ClassReadere)
}

func readConstantInfo (reader *ClassReadere, cp ConstantPool) ConstantInfo {
    tag := reader.ReadUint8()
    c := newConstantInfo(tag, cp)
    c.readInfo(reader)
    return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
    switch tag {
        case CONSTANT_Integer:  return &ConstantIntegerInfo{}
        case CONSTANT_Float:    return &ConstantFlaotInfo{}
        case CONSTANT_Long:     return &ConstantLongInfo{}
        case CONSTANT_Double:   return &ConstantDoubleInfo{}
        case CONSTANT_Utf8:     return &ConstantUtf8Info{}
        case CONSTANT_String:   return &ConstantStringInfo{}
        case CONSTANT_Fieldref: 
            return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
        case CONSTANT_Methodref:
            return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}
        case CONSTANT_InterfaceMethodref:
            return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
        case CONSTANT_NameAndType: return &ConstantNameAndTypeInfo{}
        case CONSTANT_MethodType:   return &ConstantMethodHandleInfo{}
        case CONSTANT_InvokeDynamic:   return &ConstantInvokeDynamicInfo{}
        default :   panic("java.lang.ClassFormateError: constant pool tag!")
    }
}













