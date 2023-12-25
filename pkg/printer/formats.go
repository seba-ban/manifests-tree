package printer

import (
	"fmt"
	"strings"

	"golang.org/x/exp/maps"
)

var availableFormats = map[string]PrinterFactory{
	"count": NewCountPrinter,
	"json":  NewJsonPrinter,
	"tree":  NewTreePrinter,
	"yaml":  NewYamlPrinter,
}

var AvailableFormats = maps.Keys(availableFormats)
var AvailableFormatsString = strings.Join(AvailableFormats, ", ")

func GetPrinter(format string, opts PrinterOpts) (Printer, error) {
	printer, ok := availableFormats[strings.ToLower(format)]
	if !ok {
		return nil, fmt.Errorf("unknown output format: %s, available formats are: %s", format, AvailableFormatsString)
	}
	return printer(opts), nil
}
