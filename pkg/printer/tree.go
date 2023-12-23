package printer

import (
	"fmt"

	"github.com/seba-ban/manifests-tree/pkg/store"
	"github.com/xlab/treeprint"
)

type TreePrinter struct {
	PrinterOpts
}

func (p *TreePrinter) Print(data store.TreeData) {
	tree := treeprint.New()

	var apiVersionNode treeprint.Tree
	lastApiVersion := ""
	for _, row := range getSortedPrinterRows(data) {
		if row.ApiVersion != lastApiVersion {
			apiVersionNode = tree.AddBranch(row.ApiVersion)
			lastApiVersion = row.ApiVersion
		}

		kindNode := apiVersionNode.AddBranch(fmt.Sprintf("%s [%d]", row.Kind, len(row.Documents)))
		if p.OnlyKinds {
			continue
		}

		for _, doc := range row.Documents {
			b := kindNode.AddBranch(doc.Name())
			if p.WithPaths {
				b.AddBranch(fmt.Sprintf("%s:%d", doc.Source, doc.StartLine))
			}
		}
	}

	fmt.Println(tree.String())
}

func NewTreePrinter(opts ...PrinterOption) *TreePrinter {
	printerOpts := &PrinterOpts{
		WithPaths: false,
		OnlyKinds: false,
	}

	for _, opt := range opts {
		opt(printerOpts)
	}

	return &TreePrinter{
		PrinterOpts: *printerOpts,
	}
}
