package set

import (
	"bytes"
	"fmt"
	"io"
	"path"

	"github.com/spf13/afero"
)

type image struct {
	AbsolutePath string
}

func (i image) Type() string {
	return path.Ext(i.AbsolutePath)[1:]
}

func (i image) Content(fs *afero.Afero) (io.Reader, error) {
	rawContent, err := fs.ReadFile(i.AbsolutePath)
	if err != nil {
		return nil, fmt.Errorf("reading file %s: %w", i.AbsolutePath, err)
	}

	return bytes.NewReader(rawContent), nil
}
