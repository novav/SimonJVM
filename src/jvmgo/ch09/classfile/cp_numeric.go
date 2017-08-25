package classfile

import "math"

/*  3.3.3 : CONSTANT_Integer_info
    xinxin.shi
    2017-06-03 16:56:56
*/

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
使用4字节存储整数数据
*/
type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}
func (self *ConstantIntegerInfo) Value() int32 {
	return self.val
}

/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
使用4字节存储IEEE754单精度浮点数常量
*/
type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}
func (self *ConstantFloatInfo) Value() float32 {
	return self.val
}

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
使用8字节存储整数常量
*/
type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}
func (self *ConstantLongInfo) Value() int64 {
	return self.val
}

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
使用8字节存储IEEE754 双精度浮点数
*/
type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}
func (self *ConstantDoubleInfo) Value() float64 {
	return self.val
}
