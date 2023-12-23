package printer

import (
	"sort"

	"github.com/seba-ban/manifests-tree/pkg/document"
	"github.com/seba-ban/manifests-tree/pkg/store"
	"golang.org/x/exp/maps"
)

type printerRow struct {
	ApiVersion string
	Kind       string
	Documents  []*document.YamlDocument
}

func getSortedPrinterRows(data store.TreeData) []printerRow {
	rows := make([]printerRow, 0)

	apiVersions := maps.Keys(data)
	sort.Strings(apiVersions)

	for _, apiVersion := range apiVersions {
		kinds := maps.Keys(data[apiVersion])
		sort.Strings(kinds)

		for _, kind := range kinds {
			docs := data[apiVersion][kind]
			sort.Slice(docs, func(i, j int) bool {
				return docs[i].Name() < docs[j].Name()
			})

			rows = append(rows, printerRow{
				ApiVersion: apiVersion,
				Kind:       kind,
				Documents:  docs,
			})
		}
	}

	return rows
}
