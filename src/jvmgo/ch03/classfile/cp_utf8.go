package classfile 

import "fmt"
import "unicode/utf16"
/*
    author: xinxin.shi
    2017-06-04 00:03:15
*/
/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
CONSTANT_Utf8_info常量里放的是MUTF-8编码的字符串
*/
type ConstantUtf8Info struct {
    str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
    length := uint32(reader.readUint16())
    bytes := reader.readBytes(length)
    self.str = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
    return string(bytes)
}
