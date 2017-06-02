package calssfile

import "fmt"

/** xinxin.shi 
    2017-06-02 0:13 **/

type CalssFile struct {
	//magic 	uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags		uint16
	thisClass		uint16
	superClass		uint16
	interfaces		[]uint16
	fields 			[]*MemberInfo
	methods			[]*MemberInfo
	attributes		[]AttributeInfo
}

func Parse(classData []byte) (cf *CalssFile, err error) {
	defer func() {
		for r := recover(); r != nil {
			var ok, bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &CalssFile{}
	cf.read(cr)
	return
}

func (self *CalssFile) read(reader *ClassReader) {
    self.readAndCheckMagic(reader) 
    self.readAndCheckVersion(reader)
    self.constantPool   = readConstantPool(reader)
    self.accessFlags    = reader.readerUint16()
    self.thisClass      = reader.readerUint16()
    self.superClass     = reader.readerUint16()
    self.interfaces     = reader.readerUint16s()
    this.fields     = readMembers(reader, self.constantPool)
    this.methods    = readMembers(reader, self.constantPool)
    this.attributes = readAttributes(reader, self.constantPool)
}
	
func (self *CalssFile) readAndCheckMagic (read *ClassReader) {
    magic := reader.readerUint32()
    if magic != 0xCAFEBABE {
        panic("java.lang.ClassFormatError: maigc!")
    }
}
func (self *CalssFile) readAndCheckVersion (read *ClassReader) {
    self.minorVersion = reader.readerUint16()
    self.majorVersion = reader.readerUint16()
    switch self.majorVersion {
    case 45:    // JDK 1.0.2
        return
    case 46, 47, 48, 49, 50, 51, 52:
        if (self.minorVersion == 0) {
            return
        }
    }
    panic("java.lang.UnsupportedClassVersionError!")
}
func (self *CalssFile) MinorVersion() uint16 (read *ClassReader) {
    return self.minorVersion
}
func (self *CalssFile) MajorVersion() uint16 (read *ClassReader) {
    return self.MajorVersion
}
func (self *CalssFile) ConstantPool() ConstantPool {}
func (self *CalssFile) AccessFlags() uint16 {} // getter
func (self *CalssFile) Fields() []*MemberInfo {}	// getter
func (self *CalssFile) Methods() []*MemberInfo {} //getter
func (self *CalssFile) ClassName() string {
    return self.constantPool.getClassName(self.thisClass)
}
func (self *CalssFile) SupperClass() string {
    if self.superClass > 0 {
        return self.constantPool.getClassName(self.superClass)
    }
    return ""
}
func (self *CalssFile) InterfaceNames() []string {
    interfaceNames := make([]string, len(self.interfaces))
    for i, cpIndex := range self.interfaces {
        interfaceNames[i] = self.constantPool.getClassName(cpIndex)
    }
    return interfaceNames
}