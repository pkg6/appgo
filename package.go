package appgo

import "image"

type Package interface {
	Name() string
	Version() string
	Identifier() string
	Build() string
	Channel() string
	Icon() image.Image
	Size() int64
}
