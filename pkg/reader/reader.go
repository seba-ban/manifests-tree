package reader

import (
	"errors"
	"fmt"
	"io"
	"path/filepath"

	"github.com/seba-ban/manifests-tree/pkg/document"
	"gopkg.in/yaml.v3"
)

func ReadYamlFileFromReader(reader io.ReadCloser, sourceName string) ([]*document.YamlDocument, error) {
	defer reader.Close()

	docs := make([]*document.YamlDocument, 0)
	dec := yaml.NewDecoder(reader)

	for {
		doc := &document.YamlDocument{}
		err := dec.Decode(doc)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		doc.Source = sourceName
		docs = append(docs, doc)
	}

	return docs, nil
}

func ReadYamls(dirOpts *DirWalkerOpts, inputs ...string) ([]*document.YamlDocument, error) {
	docs := make([]*document.YamlDocument, 0)

	for _, input := range inputs {

		inputType := GetInputType(input)

		if inputType != DirInputType {
			var reader io.ReadCloser
			var err error
			switch GetInputType(input) {

			case FileInputType:
				reader, err = GetFileReader(input)
				absPath, filepathErr := filepath.Abs(input)
				if filepathErr != nil {
					return nil, err
				}
				input = absPath
			case UrlInputType:
				reader, err = GetHttpReader(input)
			case StdinInputType:
				reader, err = GetStdinReader(input)
				input = "stdin"
			case UnknownInputType:
				return nil, fmt.Errorf("unknown input type: %s", input)
			}

			if err != nil {
				return nil, err
			}

			doc, err := ReadYamlFileFromReader(reader, input)
			if err != nil {
				return nil, err
			}
			docs = append(docs, doc...)
		} else {
			// handle dir
			ch, err := dirOpts.GetFilesFromDir(input)
			if err != nil {
				return nil, err
			}

			for file := range ch {
				reader, err := GetFileReader(file)
				if err != nil {
					return nil, err
				}

				doc, err := ReadYamlFileFromReader(reader, file)
				if err != nil {
					return nil, err
				}
				docs = append(docs, doc...)
			}
		}

	}

	return docs, nil
}
