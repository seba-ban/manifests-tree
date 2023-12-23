package store

import (
	"github.com/seba-ban/manifests-tree/pkg/document"
)

type KindT = string
type ApiVersionT = string
type TreeData map[ApiVersionT]map[KindT][]*document.YamlDocument

type DocumentsStore struct {
	data TreeData
}

func NewDocumentsStore() *DocumentsStore {
	store := &DocumentsStore{
		data: make(TreeData),
	}
	return store
}

func (s *DocumentsStore) Add(doc *document.YamlDocument) {
	apiVersion := doc.ApiVersion()
	kind := doc.Kind()

	apis, ok := s.data[apiVersion]

	if !ok {
		apis = make(map[KindT][]*document.YamlDocument)
		s.data[apiVersion] = apis
	}
	apis[kind] = append(apis[kind], doc)
}

func (s *DocumentsStore) AddMany(doc []*document.YamlDocument) {
	for _, d := range doc {
		s.Add(d)
	}
}

func (s *DocumentsStore) ApiVersions() []string {
	versions := make([]string, 0)
	for version := range s.data {
		if version != document.UnrecognizedApiVersion {
			versions = append(versions, version)
		}
	}
	return versions
}

func (s *DocumentsStore) Kinds(apiVersion string) []string {
	kinds := make([]string, 0)

	kindsMap, ok := s.data[apiVersion]
	if !ok {
		return kinds
	}

	for kind := range kindsMap {
		if kind != document.UnrecognizedKind {
			kinds = append(kinds, kind)
		}
	}
	return kinds
}

func (s *DocumentsStore) Documents(apiVersion string, kind string) []*document.YamlDocument {
	o, ok := s.data[apiVersion]
	if !ok {
		return make([]*document.YamlDocument, 0)
	}

	docs, ok := o[kind]
	if !ok {
		return make([]*document.YamlDocument, 0)
	}

	return docs
}

func (s *DocumentsStore) InvalidDocuments() []*document.YamlDocument {
	return s.Documents(document.UnrecognizedApiVersion, document.UnrecognizedKind)
}

func (s *DocumentsStore) TreeData() TreeData {
	return s.data
}
