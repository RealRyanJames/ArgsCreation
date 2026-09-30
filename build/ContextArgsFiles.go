package main

import (
	rsfilecontext "ArgsCreation/RSFileContext"
	tsfilecontext "ArgsCreation/TSFileContext"
	"fmt"
	"os"
)

type ArgsContexts string

type ArgsContext struct {
	args ArgsContexts
}

func (ctx ArgsContext) GetContext() string {
	return fmt.Sprintf("%s", ctx.args)
}

type LengthGet struct {
	Length int
}

func (Length LengthGet) GetLength() int {
	return Length.Length
}

func main() {

	length := LengthGet{
		Length: -1,
	}

	if length.GetLength() == -1 {

		ctx := ArgsContext{
			args: "/TS",
		}

		if os.Args[1] == ctx.GetContext() {
			tsfilecontext.TSContext()
		} else if os.Args[1] == "/RS" {
			rsfilecontext.RSFile()
		}
	}
}

type CommandWindowPrompt struct {
	CheckIsOpen func(file string) string
}
