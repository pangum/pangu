package param

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"io"
	"os"
	"reflect"

	"github.com/pangum/pangu/internal/internal/asset"
	"github.com/pangum/pangu/internal/internal/banner"
	"github.com/zs5460/art"
)

type Banner struct {
	Data any
	Type banner.Type
}

func newBanner() *Banner {
	return &Banner{
		Data: "Pangu",
		Type: banner.TypeAscii,
	}
}

func (b *Banner) Print() (err error) {
	content := ""
	switch b.Type {
	case banner.TypeTxt:
		var data []byte
		data, err = os.ReadFile(b.Data.(string))
		content = string(data)
	case banner.TypeString:
		content = b.Data.(string)
	case banner.TypeAscii:
		content = art.String(b.Data.(string))
	case banner.TypeFilepath:
		content, err = b.asciiFromFilepath(b.Data.(string))
	case banner.TypeBinary:
		content, err = b.asciiFromBytes(b.Data.([]byte))
	case banner.TypeFile:
		content, err = b.asciiFromReader(b.Data.(*os.File))
	}
	if nil != err {
		return
	}

	fmt.Println(content)
	fmt.Print(asset.Dividing)
	fmt.Println()

	return
}

func (b *Banner) asciiFromFilepath(path string) (data string, err error) {
	var imgFile *os.File

	if imgFile, err = os.Open(path); nil != err {
		return
	}
	defer func() {
		_ = imgFile.Close()
	}()
	data, err = b.asciiFromReader(imgFile)

	return
}

func (b *Banner) asciiFromBytes(fileBytes []byte) (string, error) {
	return b.asciiFromReader(bytes.NewReader(fileBytes))
}

func (b *Banner) asciiFromReader(reader io.Reader) (data string, err error) {
	var (
		conf image.Config
		img  image.Image
	)

	if conf, _, err = image.DecodeConfig(reader); nil != err {
		return
	}
	if img, _, err = image.Decode(reader); nil != err {
		return
	}

	table := []byte(banner.Ascii)
	buffer := new(bytes.Buffer)

	for i := 0; i < conf.Height; i++ {
		for j := 0; j < conf.Width; j++ {
			g := color.GrayModel.Convert(img.At(j, i))
			y := reflect.ValueOf(g).FieldByName(`Y`).Uint()
			pos := int(y * 16 / 255)
			if err = buffer.WriteByte(table[pos]); nil != err {
				return
			}
		}
		if err = buffer.WriteByte('\n'); nil != err {
			return
		}
	}
	data = buffer.String()

	return
}
