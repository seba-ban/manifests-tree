package filter

import (
	"strings"

	"github.com/seba-ban/manifests-tree/pkg/document"
	"github.com/seba-ban/manifests-tree/pkg/store"
)

type GenericFilter struct {
	IncludeKinds       []string
	IncludeKindsStrict bool
	ExcludeKinds       []string
	ExcludeKindsStrict bool

	IncludeApiVersions       []string
	IncludeApiVersionsStrict bool
	ExcludeApiVersions       []string
	ExcludeApiVersionsStrict bool

	IncludeNames       []string
	IncludeNamesStrict bool
	ExcludeNames       []string
	ExcludeNamesStrict bool

	IncludeUnrecognized bool
}

func (o *GenericFilter) matchString(source string, target string, strict bool) bool {
	if strict && source == target {
		return true
	} else if !strict && strings.Contains(target, source) {
		return true
	}
	return false
}

func (o *GenericFilter) match(list []string, strict bool, target string, returnValOnEmptyList bool) bool {
	if len(list) == 0 {
		return returnValOnEmptyList
	}

	for _, item := range list {
		if o.matchString(item, target, strict) {
			return true
		}
	}

	return false
}

func (o *GenericFilter) matchInclude(list []string, strict bool, target string) bool {
	return o.match(list, strict, target, true)
}

func (o *GenericFilter) matchExclude(list []string, strict bool, target string) bool {
	return o.match(list, strict, target, false)
}

func (o *GenericFilter) RunFilter(data store.TreeData) store.TreeData {

	filteredData := make(store.TreeData)

	for apiVersion, kinds := range data {

		if apiVersion == document.UnrecognizedApiVersion {
			continue
		}
		if !o.matchInclude(o.IncludeApiVersions, o.IncludeApiVersionsStrict, apiVersion) {
			continue
		}
		if o.matchExclude(o.ExcludeApiVersions, o.ExcludeApiVersionsStrict, apiVersion) {
			continue
		}

		for kind, docs := range kinds {

			if !o.matchInclude(o.IncludeKinds, o.IncludeKindsStrict, kind) {
				continue
			}
			if o.matchExclude(o.ExcludeKinds, o.ExcludeKindsStrict, kind) {
				continue
			}

			for _, doc := range docs {
				if !o.matchInclude(o.IncludeNames, o.IncludeNamesStrict, doc.Name()) {
					continue
				}
				if o.matchExclude(o.ExcludeNames, o.ExcludeNamesStrict, doc.Name()) {
					continue
				}

				filteredApiVersions, ok := filteredData[apiVersion]
				if !ok {
					filteredApiVersions = make(map[string][]*document.YamlDocument)
					filteredData[apiVersion] = filteredApiVersions
				}
				filteredApiVersions[kind] = append(filteredApiVersions[kind], doc)
			}
		}
	}

	if o.IncludeUnrecognized {
		filteredData[document.UnrecognizedApiVersion] = data[document.UnrecognizedApiVersion]
	}

	return filteredData
}
