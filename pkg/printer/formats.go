package printer

import (
	"fmt"
	"strings"
)

func GetPrinter(format string, opts PrinterOpts) (Printer, error) {
	switch strings.ToLower(format) {
	case "json":
		return &JsonPrinter{PrinterOpts: opts}, nil
	case "yaml":
		return &YamlPrinter{PrinterOpts: opts}, nil
	case "tree":
		return &TreePrinter{PrinterOpts: opts}, nil
	default:
		return nil, fmt.Errorf("unknown output format: %s", format)
	}
}
