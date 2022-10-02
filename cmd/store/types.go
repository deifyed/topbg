package store

import "io"

type image struct {
	Image     io.Reader
	Extension string
}
