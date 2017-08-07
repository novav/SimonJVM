package classfile

/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/

type DeprecatedAttribute struct {
    MarkerAttribute
}

type SyntheticAttribute struct {
    MarkerAttribute
}

type MarkerAttribute struct {}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
    // read nothing
}
