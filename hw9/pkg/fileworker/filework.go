package fileworker

import "io"

// WriteFile записывает только строки
func WriteFile(w io.Writer, data ...interface{}) {
	for i := range data {
		switch data[i].(type) {
		case string:
			str := data[i].(string)
			w.Write([]byte(str))
		default:
			continue
		}
	}
}
