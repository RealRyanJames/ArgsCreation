package rsfilecontext

import (
	typesparams "ArgsCreation/TypesParams"
	"fmt"
	"os"
	"strings"
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

func isNil(e error) bool {

	return e != nil
}

func RSFile() {

	cmd := CommandWindowPrompt{
		CheckIsOpen: func(fileName string) string {

			ctx := ArgsContext{
				args: ArgsContexts(strings.ToLower(fileName)),
			}

			crt := typesparams.FileType{
				RS: "main.rs",
			}

			crt_2 := typesparams.FileType{
				JS: "main.ts",
			}

			if ctx.GetContext() == crt.GetFileType() {
				fileName = crt.GetFileType()
			} else if ctx.GetContext() == crt.GetFileType() {
				fileName = crt_2.GetFileType()
			}

			f, err := os.Create(ctx.GetContext())
			if isNil(err) {
				fmt.Println(err)
			}

			defer f.Close()
			fmt.Println(f.Name())

			return f.Name()
		},
	}

	ctx_rs := ArgsContext{
		args: "/RS",
	}

	if os.Args[1] == ctx_rs.GetContext() {
		cmd.CheckIsOpen("main.rs")
	}

}
