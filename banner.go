package pangu

import (
	`fmt`
	`io/ioutil`
)

const (
	// BannerTypeFile 文件
	BannerTypeFile BannerType = "file"
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
		if data, err = ioutil.ReadFile(b.content); nil != err {
			return
		}
		content = string(data)
	}
	fmt.Println(content)
	fmt.Print("\n\n")

	return
}
