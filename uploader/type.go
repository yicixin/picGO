package uploader

import "io"

type Uploader interface {
	Upload(filename string, reader io.Reader) (string, error)
}
