package tools

import "time"

type ConvertType interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 | time.Time
}

func ToPointer[T ConvertType](data T) (resp *T) {
	resp = &data
	return
}
