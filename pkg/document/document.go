package document

import (
	"encoding/json"
	"strings"

	"github.com/kubernetes-client/go/kubernetes/client"
	"gopkg.in/yaml.v3"
)

var UnrecognizedKind = "unrecognized"
var UnrecognizedApiVersion = "unrecognized"
var NameUnknown = "noname"

type YamlDocument struct {
	StartLine  int
	Source     string
	apiVersion string
	kind       string
	meta       *client.V1ObjectMeta
	metaValid  bool
}

func (d *YamlDocument) UnmarshalYAML(value *yaml.Node) error {

	d.StartLine = value.Line

	data := make(map[string]interface{})
	err := value.Decode(&data)
	if err != nil {
		data = make(map[string]interface{})
	}

	// get metadata
	meta := &client.V1ObjectMeta{}
	jsonMeta, ok := data["metadata"]
	if ok {
		rawMeta, err := json.Marshal(jsonMeta)
		if err == nil {
			err = json.Unmarshal(rawMeta, meta)
			if err == nil {
				d.metaValid = true
			}
		}
	}
	d.meta = meta

	// get apiVersion
	apiVersion, ok := data["apiVersion"].(string)
	if !ok {
		apiVersion = UnrecognizedApiVersion
	}
	d.apiVersion = strings.ToLower(strings.TrimSpace(apiVersion))

	// get kind
	kind, ok := data["kind"].(string)
	if !ok {
		kind = UnrecognizedKind
	}
	d.kind = strings.ToLower(strings.TrimSpace(kind))

	return nil
}

func (d *YamlDocument) IsValid() bool {
	return d.metaValid && d.Kind() != UnrecognizedKind && d.ApiVersion() != UnrecognizedKind
}

func (d *YamlDocument) Kind() string {
	return d.kind
}

func (d *YamlDocument) ApiVersion() string {
	return d.apiVersion
}

func (d *YamlDocument) Name() string {
	if !d.metaValid || d.meta.Name == "" {
		return NameUnknown
	}
	return d.meta.Name
}
