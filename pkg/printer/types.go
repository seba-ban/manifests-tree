package printer

import "github.com/seba-ban/manifests-tree/pkg/store"

type Printer interface {
	Print(store.TreeData)
}

type PrinterOpts struct {
	WithPaths bool
	OnlyKinds bool
}

type PrinterOption func(*PrinterOpts)

func WithPaths(withPaths bool) PrinterOption {
	return func(o *PrinterOpts) {
		o.WithPaths = withPaths
	}
}

func WithOnlyKinds(onlyKinds bool) PrinterOption {
	return func(o *PrinterOpts) {
		o.OnlyKinds = onlyKinds
	}
}
