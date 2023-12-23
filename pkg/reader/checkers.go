package reader

import (
	"os"
	"strings"
)

type InputType int

const (
	UrlInputType InputType = iota
	FileInputType
	StdinInputType
	DirInputType
	UnknownInputType
)

func GetInputType(input string) InputType {
	// stdin
	if input == "-" {
		return StdinInputType
	}

	// url
	if strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://") {
		return UrlInputType
	}

	// file and dir
	if info, err := os.Stat(input); err == nil {
		mode := info.Mode()

		if mode.IsRegular() {
			return FileInputType
		}

		if mode.IsDir() {
			return DirInputType
		}
	}

	return UnknownInputType
}
