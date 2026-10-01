package typesparams

import "reflect"

const (
	RS = 0
	JS = 1
)

type FileType struct {
	RS    string
	JS    string
	other string
}

func (fileType FileType) GetFileType() string {
	if reflect.TypeOf(fileType.JS).Kind().String() == "string" {
		return fileType.JS
	} else if reflect.TypeOf(fileType.RS).Kind().String() == "string" {
		return fileType.RS
	}

	return string(fileType.other)
}
