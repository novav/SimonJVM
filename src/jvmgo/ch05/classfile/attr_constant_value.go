package classfile

/* ConstantValue 定长属性 只用与field_info,  用于表示常量表达式的值 */

/*
ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_lenth;
    u2 constantvalue_index;
}
*/

type ConstantValueAttribute struct {
    constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
    self.constantValueIndex = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
    return self.constantValueIndex
}
