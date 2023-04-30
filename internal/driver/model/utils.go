package model

type Util interface {
	Get(n int) []byte
	Read(buf []byte) error
	FillMeta()
}
