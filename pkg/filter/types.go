package filter

import "github.com/seba-ban/manifests-tree/pkg/store"

type Filter interface {
	RunFilter(store.TreeData) store.TreeData
}
