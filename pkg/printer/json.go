package printer

import (
	"encoding/json"
	"fmt"

	"github.com/seba-ban/manifests-tree/pkg/store"
)

type JsonPrinter struct {
	PrinterOpts
}

func (p *JsonPrinter) GetMap(data store.TreeData) map[string]map[string][]map[string]interface{} {
	printData := make(map[string]map[string][]map[string]interface{})

	for _, row := range getSortedPrinterRows(data) {

		if _, ok := printData[row.ApiVersion]; !ok {
			printData[row.ApiVersion] = make(map[string][]map[string]interface{})
		}

		if _, ok := printData[row.ApiVersion][row.Kind]; !ok {
			printData[row.ApiVersion][row.Kind] = make([]map[string]interface{}, 0)
		}

		if p.OnlyKinds {
			continue
		}

		for _, doc := range row.Documents {

			o := map[string]interface{}{
				"name": doc.Name(),
			}

			if p.WithPaths {
				o["source"] = doc.Source
				o["start_line"] = doc.StartLine
			}

			printData[row.ApiVersion][row.Kind] = append(printData[row.ApiVersion][row.Kind], o)
		}
	}

	return printData
}

func (p *JsonPrinter) Print(data store.TreeData) {
	printData := p.GetMap(data)

	dumped, err := json.MarshalIndent(printData, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(dumped))
}

func NewJsonPrinter(opts ...PrinterOption) *JsonPrinter {
	printerOpts := &PrinterOpts{
		WithPaths: false,
		OnlyKinds: false,
	}

	for _, opt := range opts {
		opt(printerOpts)
	}

	return &JsonPrinter{
		PrinterOpts: *printerOpts,
	}
}
