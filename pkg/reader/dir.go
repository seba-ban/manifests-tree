package reader

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type DirWalkerOpts struct {
	Recursive bool
}

// GetFilesFromDir returns a channel of files found in the given directory.
// If recursive is true, it will also include files in subdirectories.
func (d *DirWalkerOpts) GetFilesFromDir(dir string) (<-chan string, error) {

	dir, err := filepath.Abs(dir)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	channel := make(chan string)

	go (func(dir string, recursive bool, channel chan string) {
		defer close(channel)

		fileSystem := os.DirFS(dir)
		fs.WalkDir(fileSystem, ".", func(p string, d fs.DirEntry, err error) error {
			if err != nil {
				log.Fatal(err)
				return nil
			}

			if d.Type().IsDir() {
				if p != "." && !recursive {
					return fs.SkipDir
				}
				return nil
			}

			ext := filepath.Ext(p)
			if ext == ".yaml" || ext == ".yml" {
				channel <- filepath.Join(dir, p)
			}

			return nil
		})
	})(dir, d.Recursive, channel)

	return channel, nil
}
