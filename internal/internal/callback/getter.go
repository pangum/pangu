package callback

type Getter interface {
	Get(key string) string
}
