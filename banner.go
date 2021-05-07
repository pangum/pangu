package pangu

import (
	`bytes`
	`fmt`
	`image`
	`image/color`
	`io/ioutil`
	`os`
	`reflect`
)

const (
	// BannerTypeTxt 文本文件
	BannerTypeTxt BannerType = "txt"
	// BannerTypeFilepath 图片文件
	BannerTypeFilepath BannerType = "filepath"
	// BannerTypeString 直接显示
	BannerTypeString BannerType = "string"
	// BannerTypeAscii 内部转换
	BannerTypeAscii BannerType = "ascii"
	// BannerTypeBinary 二进制文件数据
	BannerTypeBinary BannerType = "binary"
	// BannerTypeFile 文件数据
	BannerTypeFile BannerType = "file"

	ascii = "MND8OZ$7I?+=~:,.."
)

type (
	// BannerType 标志类型
	BannerType string

	banner struct {
		content    string
		bannerType BannerType
	}
)

func (b *banner) print() (err error) {
	var content string

	switch b.bannerType {
	case BannerTypeTxt:
		var data []byte
		data, err = ioutil.ReadFile(b.content)
		content = string(data)
	case BannerTypeString:
		content = b.content
	case BannerTypeAscii:
		content = b.content
	case BannerTypeFilepath:
		content, err = b.convertToAscii(b.content)
	case BannerTypeBinary:
	case BannerTypeFile:
	}
	if nil != err {
		return
	}

	fmt.Println(content)
	fmt.Print("\n\n")

	return
}

func (b *banner) convertToAscii(path string) (data string, err error) {
	var (
		imgFile *os.File
		conf    image.Config
		img     image.Image
	)

	if imgFile, err = os.Open(path); nil != err {
		return
	}

	if conf, _, err = image.DecodeConfig(imgFile); nil != err {
		return
	}
	if img, _, err = image.Decode(imgFile); nil != err {
		return
	}
	defer func() {
		_ = imgFile.Close()
	}()

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
