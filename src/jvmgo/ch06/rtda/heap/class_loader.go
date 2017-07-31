package heap
import "fmt"
import "jvmgo/ch06/classfile"
import "jvmgo/ch06/classpath"

/*
    6.3
    1、find & read class file 
    2、parse file (-> 生成JVM可用的类数据)
    3、link 
    xinxin.shi
    2017-07-30 21:02:06
*/
type ClassLoader struct {
    cp          *classpath.Classpath
    classMap    map[string]*Class // loaded classes
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
    return &ClassLoader {
        cp : cp,
        classMap: make(map[string]*Class),
    }
}

func (self *ClassLoader) LoadClass(name string) *Class {
    if class, ok := self.classMap[name]; ok {
        return class // 类已经加载
    }
    return self.loadNonArrayClass(name)
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
    data, entry := self.readClass(name)
    class := self.defineClass(data)
    link(class)
    fmt.Printf("[Loaded %s from %s]\n", name, entry)
    return class
}

func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
    data, entry, err := self.cp.ReadClass(name);
    if err != nil {
        panic("java.lang.ClassNotFountException: " + name)
    }
    return data, entry
}
/*
    6.3.2
*/
func (self *ClassLoader) defineClass(data []byte) *Class {
    class := parseClass(data)
    class.loader = self
    resolvedSuperClass(class)
    resolvedInterfaces(class)
    self.classMap[class.name] = class
    return class
}

func parseClass(data []byte) *Class{
    cf, err := classfile.Parse(data)
    if err != nil {
        panic("java.lang.ClassFormateError")
    }
    return newClass(cf) // 6.1.1
}

func resolvedSuperClass(class *Class) {
    if class.name != "java/lang/Object" {
        class.superClass = class.loader.LoadClass(class.superClassName)
    }
}


func resolvedInterfaces(class *Class) {
    interfaceCount := len(class.interfaceNames)
    if interfaceCount > 0 {
        class.interfaces = make([]*Class, interfaceCount)
        for i, interfaceName := range class.interfaceNames {
            class.interfaces[i] = class.loader.LoadClass(interfaceName)
        }

    }
}

func link(class *Class) {
    verify(class)
    prepare(class)
}

func verify(class *Class) {
    // todo
}

/*chap6.4*/
func prepare(class *Class) {
    calcInstanceFieldSlotIds(class)
    calcStaticFieldSlotIds(class)
    allocAndInitStaticVars(class)
}

/*计算实例字段的个数，同时给它们编号*/
func calcInstanceFieldSlotIds(class *Class) {
    slotId := uint(0)
    if class.superClass != nil {
        slotId = class.superClass.instanceSlotCount
    }
    for _, field := range class.fields {
        if !field.IsStatic() {
            field.slotId = slotId
            slotId ++
            if field.isLongOrDouble() {
                slotId ++
            }
        }
    }
    class.instanceSlotCount = slotId
}

/*计算静态字段的个数，同时给它们编号*/
func calcStaticFieldSlotIds(class *Class) {
    slotId := uint(0)
    for _, field := range class.fields {
        if field.IsStatic() {
            field.slotId = slotId
            slotId ++
            if field.isLongOrDouble() {
                slotId++
            }
        }
    }
    class.staticSlotCount  = slotId
}

/*给类变量分配空间，然后给它们赋予初始值*/
func allocAndInitStaticVars(class *Class) {
    class.staticVars = newSlots(class.staticSlotCount)
    for _, field := range class.fields {
        if field.IsStatic() && field IsFinal() {
            initStaticFinalVar(class, field)
        }
    }
}

func initStaticFinalVar(class *Class, field *Field) {
    vars := class.staticVars
    cp := class.constantPool
    cpIndex := field.ConstValueIndex()
    slotId := field.SlotId()
    if cpIndex > 0 {
        switch field.Descriptor() {
            case "Z", "B", "C", "S", "I":
                val := cp.GetConstant(cpIndex).(int32)
                vars.setInt(slotId, val)
            case "J":
                val := cp.GetConstant(cpIndex).(int64)
                vars.SetLong(slotId, val)
            case "F":
                val := cp.GetConstant(cpIndex).(float32)
                vars.SetFloat(slotId, val)
            case "D":
                val := cp.GetConstant(cpIndex).(float64)
                vars.SetDouble(slotId, val)
            case "Ljava/lang/String" :
                panic("todo") // chap 8
        }
    }
}