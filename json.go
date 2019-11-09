package gons3

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
)

func jsonUnmarshal(reader io.Reader, v interface{}) error {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, v)
	if err != nil {
		return err
	}

	return nil
}

func jsonMarshal(data interface{}) jsonReader {
	b, err := json.Marshal(data)
	if err != nil {
		return jsonReader{Error: err}
	}
	return jsonReader{Reader: bytes.NewReader(b)}
}

type jsonReader struct {
	Error  error
	Reader io.Reader
}

func (j jsonReader) Read(p []byte) (n int, err error) {
	if j.Error != nil {
		return 0, j.Error
	}
	return j.Reader.Read(p)
}
