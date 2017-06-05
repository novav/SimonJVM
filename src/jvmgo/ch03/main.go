package main

import "fmt"
import "strings"
import "jvmgo/ch03/classfile"
import "jvmgo/ch03/classpath"

	/*
	xinxin.shi
	2017-06-04 23:45:47
	*/
	/*
	介绍了Java虚拟机从哪里搜索class文件，并且实现了类路径功能，已经可以把class文件读取到内存中
	*/

	func main() {
		cmd := parseCmd()

		if cmd.versionFlag {
			fmt.Println("version 0.0.1")
		} else if cmd.helpFlag || cmd.class == "" {
			printUsage()
		} else {
			startJVM(cmd)
		}
	}
	
	func startJVM(cmd *Cmd) {
		// ch02.exe -Xjre D:\Tools\Java\jdk1.8.0_77\jre\ 	java.lang.Object
		//				  cmd.XjreOption					cmd.cpOption
		fmt.Printf("main->startJVM() => cmd.XjreOption:%v, cmd.cpOption:%v \n", cmd.XjreOption, cmd.cpOption)
		// cmd.XjreOption:D:\Tools\Java\jdk1.8.0_77\jre\, cmd.cpOption:cmd.XjreOption:D:\Tools\Java\jdk1.8.0_77\jre\, cmd.cpOption:

		cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)

		fmt.Printf("main->startJVM() => classpath:%v class:%v args:%v \n", cp, cmd.class, cmd.args)
		// classpath:D:\Tools\go\workspace\bin class:java.lang.Object args:[]

		className := strings.Replace(cmd.class, ".", "/", -1)

		cf := loadClass(className, cp)
		fmt.Println(cmd.class)
		printClassInfo(cf)

	}

	func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
		classData, _, err := cp.ReadClass(className)
		if err != nil {
			panic(err)
		}
		cf, err := classfile.Parse(classData)
		if err != nil {
			panic(err)
		}
		return cf
	}

	func printClassInfo(cf *classfile.ClassFile) {
		fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
		fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
		fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
		fmt.Printf("this class: %v\n", cf.ClassName())
		fmt.Printf("super class: %v\n", cf.SuperClassName())
		fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
		fmt.Printf("fields count: %v\n", len(cf.Fields()))
		for _, f := range cf.Fields() {
			fmt.Printf("	%s\n", f.Name())
		}
        
	    fmt.Printf("methods count: %v\n", len(cf.Methods()))
		for _, m := range cf.Methods() {
			fmt.Printf("	%s\n", m.Name())
		}
	}