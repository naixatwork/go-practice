package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	quitch := make(chan struct{})
	js.Global().Set("sayHello", js.FuncOf(sayHello))
	js.Global().Set("add", js.FuncOf(add))
	fmt.Println("hello!!!")

	select {}

	<-quitch
}

func add(val js.Value, values []js.Value) any {
	fmt.Printf("%+v\n", val)
	fmt.Printf("%+v\n", values)
	return nil
}

func sayHello(val js.Value, values []js.Value) any {
	fmt.Println("hello from golang!!!!")
	return nil
}
