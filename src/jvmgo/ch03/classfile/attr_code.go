package classfile

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
    3.4.5 
    xinxin.shi
    2017-06-04 22:59:11
*/

type CodeAttribute struct {
    cp          ConstantPool
    maxStack    uint16
    maxLocals   uint16
    code        []byte
    exceptionTable []*ExceptionTableEntry
    attributes  []AttributeInfo
}    

func (self *CodeAttribute) readInfo(reader *ClassReader) {
    self.maxStack = reader.readUint16()
    self.maxLocals = reader.readUint16()
    codeLength := reader.readUint32()
    self.code = reader.readBytes(codeLength)
    self.exceptionTable = readExceptionTable(reader)
    self.attributes = readAttributes(reader, self.cp)
}


type ExceptionTableEntry struct {
    startPc     uint16
    endPc       uint16
    handlerPc   uint16
    catchType   uint16
}
func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
    exceptionTableLength := reader.readUint16()
    exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
    for i := range exceptionTable {
        exceptionTable[i] = &ExceptionTableEntry{
            startPc:    reader.readUint16(),
            endPc:      reader.readUint16(),
            handlerPc:   reader.readUint16(),
            catchType:  reader.readUint16(),
        }
    }
    return exceptionTable
}


/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;      // 最大深度
    u2 max_locals;      // 局部变量表大小
    u4 code_length;     // 字节码
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];  //异常处理表
    u2 attributes_count;
    attribute_info attributes[attributes_count];  //属性表
}
*/
