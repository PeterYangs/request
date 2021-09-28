package request

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
)

type body struct {
	body   io.ReadCloser
	header http.Header
}

func (b body) Content() (content, error) {

	defer b.body.Close()

	var err error

	var read io.ReadCloser

	read = b.body

	if b.header.Get("Content-Encoding") == "gzip" {

		read, err = gzip.NewReader(b.body)

		if err != nil {

			return content{content: []byte{}}, err
		}

	}

	bb, err := ioutil.ReadAll(read)

	if err != nil {

		return content{content: []byte{}}, err
	}

	return content{content: bb}, nil
}
