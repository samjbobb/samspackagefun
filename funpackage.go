package samspackagefun

import "fmt"

func Yell() {
	fmt.Println("HELLO FROM A SHARED FUNCTION!")
	notAnExportedFunc()
}

func notAnExportedFunc(){
	fmt.Println("Hi, I'm not exported")
}
