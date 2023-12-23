package printer

import (
	"fmt"
	"log"

	"github.com/seba-ban/manifests-tree/pkg/store"
	"gopkg.in/yaml.v3"
)

type YamlPrinter struct {
	PrinterOpts
}

func (p *YamlPrinter) Print(data store.TreeData) {
	jsonPrinter := JsonPrinter{PrinterOpts: p.PrinterOpts}
	printData := jsonPrinter.GetMap(data)
	dumped, err := yaml.Marshal(printData)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(dumped))
}

func NewYamlPrinter(opts ...PrinterOption) *YamlPrinter {
	printerOpts := &PrinterOpts{
		WithPaths: false,
		OnlyKinds: false,
	}

	for _, opt := range opts {
		opt(printerOpts)
	}

	return &YamlPrinter{
		PrinterOpts: *printerOpts,
	}
}
