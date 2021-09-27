package request

import (
	"io"
	"io/ioutil"
)

type body struct {
	body io.ReadCloser
}

func (b body) Content() (content, error) {

	bb, err := ioutil.ReadAll(b.body)

	if err != nil {

		return content{content: []byte{}}, err
	}

	defer b.body.Close()

	return content{content: bb}, nil
}
