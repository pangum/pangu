package banner

const (
	TypeUnknown Type = iota
	TypeTxt
	TypeFilepath
	TypeString
	TypeAscii
	TypeBinary
	TypeFile
)

type Type uint8
