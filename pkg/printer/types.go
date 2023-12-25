package printer

import "github.com/seba-ban/manifests-tree/pkg/store"

type Printer interface {
	Print(store.TreeData)
}

type PrinterOpts struct {
	WithPaths bool
	OnlyKinds bool
}

type PrinterFactory func(PrinterOpts) Printer
