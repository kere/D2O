package upload

import (
	"time"

	"github.com/valyala/fasthttp"
)

// Image class
type Image struct {
}

// NewImage class
func NewImage() *Image {
	return &Image{}
}

// Auth a
func (m *Image) Auth(ctx *fasthttp.RequestCtx) error {
	return nil
}

// Success f
func (m *Image) Success(name, folder string) error {
	return nil
}

// StoreDir f
func (m *Image) StoreDir() string {
	s := time.Now().Format("200601")
	return "webroot/upload/images/" + s
}
