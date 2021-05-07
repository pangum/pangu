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
	// BannerTypeFile 文件
	BannerTypeFile BannerType = "file"
	// BannerTypePicture 图片
	BannerTypePicture BannerType = "picture"
	// BannerTypeString 直接显示
	BannerTypeString BannerType = "string"
	// BannerTypeConvert 内部转换
	BannerTypeConvert BannerType = "convert"

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
	case BannerTypeFile:
		var data []byte
		data, err = ioutil.ReadFile(b.content)
		content = string(data)
	case BannerTypeString:
		content = b.content
	case BannerTypeConvert:
		content = b.content
	case BannerTypePicture:
		content, err = b.convertToAscii(b.content)
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
