package main

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"os"
)

type People struct {
	id int
}

func (p *People) GetID() int {
	fmt.Printf("calling People.GetID()\n")
	return 0
}

func HookGetID(p *People) int {
	//fmt.Fprintln(os.Stdout, "calling HookGetID")
	fmt.Printf("calling Hook People.GetID()\n")
	return TrampHookGetID(p)
}

func TrampHookGetID(p *People) int {
	// a dummy function to make room for a shadow copy of the original function.
	// it doesn't matter what we do here, just to create an addressable function with adequate size.
	TrampHookGetID(p)
	for {
		fmt.Printf("hello")
	}
	return 0
}

func myPrintln(a ...interface{}) (n int, err error) {
	fmt.Fprintln(os.Stdout, "before real Printfln")
	return myPrintlnTramp(a...)
}

func myPrintlnTramp(a ...interface{}) (n int, err error) {
	// a dummy function to make room for a shadow copy of the original function.
	// it doesn't matter what we do here, just to create an addressable function with adequate size.
	myPrintlnTramp(a...)
	myPrintlnTramp(a...)
	myPrintlnTramp(a...)

	for {
		fmt.Printf("hello")
	}

	return 0, nil
}

func main() {
	gohook.Hook(fmt.Println, myPrintln, myPrintlnTramp)
	fmt.Println("hello world!")

	p := &People{1}
	err := gohook.Hook((*People).GetID, HookGetID, TrampHookGetID)
	if err != nil {
		fmt.Println(err)
	}
	p.GetID()
}
