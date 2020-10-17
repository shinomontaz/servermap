package i8s

import (
	"fmt"
	"net/http"
)

const TypeWARNINIG string = "warning"
const TypeSUCCESS string = "success"
const TypeDANGER string = "danger"

type IErrorHandler interface {
	HandleError(w http.ResponseWriter, err error)
	Set401(w http.ResponseWriter, str string)
}

func ByteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}
