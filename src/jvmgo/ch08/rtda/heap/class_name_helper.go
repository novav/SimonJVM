package heap

/*
    Chap8.3.2
    anewarray指令
    xinxin.shi
    2017-08-18 21:32:01
*/
func getArrayClassName(className string) string {
    return "[" + toDescriptor(className)
}

func toDescriptor(className string) string {
    if className[0] == "[" {
        return className
    }
    if d, ok := primitiveTypes[className];  ok {
        return d
    }
    return "L" + className + ";"
}

var primitiveTypes = map[string] string {
    "void":         "V",
    "boolean":      "Z",    // 
    "byte":         "B",
    "short":        "S",
    "int":          "I",    
    "long":         "J",    // 
    "char":         "C",
    "float":        "F",
    "double":       "D",
}