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

func prepare(class *Class) {
    
}