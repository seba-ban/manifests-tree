package reader

import (
	"io"
	"net/http"
	"os"
)

type ReaderGetter func(string) (io.ReadCloser, error)

func GetHttpReader(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func GetFileReader(filepath string) (io.ReadCloser, error) {
	return os.Open(filepath)
}

func GetStdinReader(_ string) (io.ReadCloser, error) {
	return os.Stdin, nil
}
