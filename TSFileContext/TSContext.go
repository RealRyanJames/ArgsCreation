package tsfilecontext

import (
	"fmt"
	"os"
)

type CommandWindowPrompt struct {
	CheckIsOpen func(file string) string
}

type ArgsContexts string

type ArgsContext struct {
	args ArgsContexts
}

func (ctx ArgsContext) GetContext() string {
	return fmt.Sprintf("%s", ctx.args)
}

func TSContext() {

	cmd := CommandWindowPrompt{
		CheckIsOpen: func(fileName string) string {

			if fileName == "ts" {
				fileName = "main.ts"
			} else {
				fileName = "main.rs"
			}

			f, err := os.Create(fileName)
			if err != nil {
				fmt.Println(err)
			}

			defer f.Close()
			fmt.Println(f.Name())

			return f.Name()

		},
	}

	cmd.CheckIsOpen("ts")
	fmt.Println("TYPESCRIPT FILE")
}
