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
    CONSTANT_InvokeDynameic         = 18
}

type ConstantInfo interface {
    readInfo(reader *ClassReadere)
}

func readConstantInfo (reader *ClassReadere, cp ConstantPool) ConstantInfo {

}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
    
}
