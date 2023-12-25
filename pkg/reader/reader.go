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

		// it seems that in yaml files with lines like
		// ---
		// ---
		// there will be empty documents that won't call
		// the unmarshal function, so we need to skip them
		if doc.StartLine == 0 {
			continue
		}

		doc.Source = sourceName
		docs = append(docs, doc)
	}

	return docs, nil
}

func ReadYamls(dirOpts *DirWalkerOpts, inputs ...string) ([]*document.YamlDocument, error) {

	inputTypes := make([]InputType, len(inputs))
	seen := make(map[string]bool)
	for i, input := range inputs {
		// check for duplicates
		if _, ok := seen[input]; ok {
			return nil, fmt.Errorf("duplicate input: %s", input)
		}
		seen[input] = true

		// check if input type valid
		inputTypes[i] = GetInputType(input)
		if inputTypes[i] == UnknownInputType {
			return nil, fmt.Errorf("unknown input type: %s", input)
		}
	}

	docs := make([]*document.YamlDocument, 0)

	for i, input := range inputs {

		inputType := inputTypes[i]

		if inputType != DirInputType {
			var reader io.ReadCloser
			var err error

			switch inputType {
			case FileInputType:
				reader, err = GetFileReader(input)
				absPath, filepathErr := filepath.Abs(input)
				if filepathErr != nil {
					if err != nil {
						return nil, err
					}
					reader.Close()
					return nil, filepathErr
				}
				input = absPath
			case UrlInputType:
				reader, err = GetHttpReader(input)
			case StdinInputType:
				reader, err = GetStdinReader(input)
				input = "stdin"
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
