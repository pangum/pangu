package constraint

type Argument interface {
	~int | ~int64 |
		~uint | ~uint64 |
		~float32 | ~float64 |
		~bool |
		~string | []string
}
