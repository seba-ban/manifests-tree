package printer

import (
	"fmt"

	"github.com/seba-ban/manifests-tree/pkg/store"
)

type CountPrinter struct {
	PrinterOpts
}

func getRowFirstCol(row *printerRow) string {
	return fmt.Sprintf("%s/%s", row.ApiVersion, row.Kind)
}

func (p *CountPrinter) Print(data store.TreeData) {
	rows := getSortedPrinterRows(data)
	longest := 0
	for _, row := range rows {
		colLen := len(getRowFirstCol(&row))
		if colLen > longest {
			longest = colLen
		}
	}

	for _, row := range rows {
		fmt.Printf("%-*s  %d\n", longest, getRowFirstCol(&row), len(row.Documents))
	}
}

func NewCountPrinter(opts PrinterOpts) Printer {
	return &CountPrinter{
		PrinterOpts: opts,
	}
}
