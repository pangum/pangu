package internal

import (
	"strings"

	"github.com/goexl/gox"
	"github.com/goexl/structer"
	"github.com/pangum/pangu/internal/runtime"
)

type Decoder struct {
	from *map[string]any
}

func NewDecoder(from *map[string]any) *Decoder {
	return &Decoder{
		from: from,
	}
}

func (d *Decoder) Decode(target runtime.Pointer) error {
	return structer.Copy().From(d.from).To(target).Mapper(d.mapper).Build().Apply()
}

func (d *Decoder) mapper(key string, field string) (mapped bool) {
	if key == field {
		mapped = true
	} else {
		mapped = d.variants(key, field)
	}

	return
}

func (d *Decoder) variants(key string, field string) (mapped bool) {
	from := gox.String(field).Switch()
	checked := strings.ToLower(key)
	if strings.ToLower(from.Camel().Build().Case()) == checked {
		mapped = true
	} else if strings.ToLower(from.Underscore().Build().Case()) == checked {
		mapped = true
	} else if strings.ToLower(from.Strike().Build().Case()) == checked {
		mapped = true
	}

	return
}
