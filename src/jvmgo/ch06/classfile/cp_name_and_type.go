package classfile

/*
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
*/
/*
1）类型描述符。
    基本类型描述符 单个字母
        byte、short、char、int、long、float、double
        B、   S、    C、   I、  J、   F、    D

    引用类型的描述符是L＋类的完全限定名＋分号。
    数组类型的描述符是[＋数组元素类型描述符。
2) 字段描述符就是字段类型的描述符。
3）方法描述符是（分号分隔的参数类型描述符）+返回值类型描
述符，其中void返回值由单个字母V表示。
*/

type ConstantNameAndTypeInfo struct {
    nameIndex       uint16
    descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
    self.nameIndex = reader.readUint16()
    self.descriptorIndex = reader.readUint16()
}