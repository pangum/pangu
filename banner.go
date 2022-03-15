package pangu

import (
	`bytes`
	_ `embed`
	`fmt`
	`image`
	`image/color`
	`io`
	`io/ioutil`
	`os`
	`reflect`

	`github.com/common-nighthawk/go-figure`
)

const (
	// BannerTypeTxt 文本文件
	BannerTypeTxt bannerType = "txt"
	// BannerTypeFilepath 图片文件
	BannerTypeFilepath bannerType = "path"
	// BannerTypeString 直接显示
	BannerTypeString bannerType = "string"
	// BannerTypeAscii 内部转换
	BannerTypeAscii bannerType = "ascii"
	// BannerTypeBinary 二进制文件数据
	BannerTypeBinary bannerType = "binary"
	// BannerTypeFile 文件数据
	BannerTypeFile bannerType = "file"

	ascii = "MND8OZ$7I?+=~:,.."
)

//go:embed asset/dividing_line.txt
var dividingLine string

type (
	bannerType string

	banner struct {
		data interface{}
		typ  bannerType
	}
)

func (b *banner) print() (err error) {
	var content string

	switch b.typ {
	case BannerTypeTxt:
		var data []byte
		data, err = ioutil.ReadFile(b.data.(string))
		content = string(data)
	case BannerTypeString:
		content = b.data.(string)
	case BannerTypeAscii:
		content = figure.NewFigure(b.data.(string), "", true).String()
	case BannerTypeFilepath:
		content, err = b.asciiFromFilepath(b.data.(string))
	case BannerTypeBinary:
		content, err = b.asciiFromBytes(b.data.([]byte))
	case BannerTypeFile:
		content, err = b.asciiFromReader(b.data.(*os.File))
	}
	if nil != err {
		return
	}

	fmt.Print(content)
	fmt.Print(dividingLine)
	fmt.Print("\n")

	return
}

func (b *banner) asciiFromFilepath(path string) (data string, err error) {
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

func (b *banner) asciiFromBytes(fileBytes []byte) (string, error) {
	return b.asciiFromReader(bytes.NewReader(fileBytes))
}

func (b *banner) asciiFromReader(reader io.Reader) (data string, err error) {
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

	table := []byte(ascii)
	buffer := new(bytes.Buffer)

	for i := 0; i < conf.Height; i++ {
		for j := 0; j < conf.Width; j++ {
			g := color.GrayModel.Convert(img.At(j, i))
			y := reflect.ValueOf(g).FieldByName("Y").Uint()
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
