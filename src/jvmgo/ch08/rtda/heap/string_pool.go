package heap

import "unicode/utf16"

/*
    Chap 8.5.1
    字符串池
    xinxin.shi
    2017-08-22 23:36:00
*/

var internedStrings = map[string]*Object{}

func JString(loader *ClassLoader, goStr string) *Object {
    if internedStr, ok := stringToUtf16(goStr); ok {
        return internedStr
    }
    chars := stringToUtf16(goStr)
    jChars := &Object{loader.LoadClass("[C"), chars}
    jStr := loader.LoadClass("java/lang/String").NewObject()
    jStr.SetRefVar("value", "[C", jChars)
    internedStrings[goStr] = jStr
    return jStr
}

func stringToUtf16(s string) []uint16 {
    runes := []rune(s)
    return utf16.Encode(runes)
}