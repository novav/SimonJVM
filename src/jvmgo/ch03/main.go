package main 
	import "fmt"
	import "strings"
	import "jvmgo/ch03/classpath"

	/*
	介绍了Java虚拟机从哪里搜索class文件，并且实现了类路径功能，已经可以把class文件读取到内存中
	*/

	func main() {
		cmd := parseCmd();
		if cmd.versionFlag {
			fmt.Printf("version 0.0.1")
		} else if cmd.helpFlag || cmd.class == "" {
			printUsage()
		} else {
			startJVM(cmd)
		}
	}
	
	func startJVM(cmd *Cmd) {
		// ch02.exe -Xjre D:\Tools\Java\jdk1.8.0_77\jre\ 	java.lang.Object
		//				  cmd.XjreOption					cmd.cpOption
		fmt.Printf("cmd.XjreOption:%v, cmd.cpOption:%v \n", cmd.XjreOption, cmd.cpOption)
		// cmd.XjreOption:D:\Tools\Java\jdk1.8.0_77\jre\, cmd.cpOption:cmd.XjreOption:D:\Tools\Java\jdk1.8.0_77\jre\, cmd.cpOption:

		cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)

		fmt.Printf("classpath:%v class:%v args:%v \n", cp, cmd.class, cmd.args)
		// classpath:D:\Tools\go\workspace\bin class:java.lang.Object args:[]

		className := strings.Replace(cmd.class, ".", "/", -1)
		classData, _, err := cp.ReadClass(className)
		if err != nil {
			fmt.Printf("Could not find or load main class %s\n " , cmd.class)
		}
		fmt.Printf("class data:%v\n", classData)
	}