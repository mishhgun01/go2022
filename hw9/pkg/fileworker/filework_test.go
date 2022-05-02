package fileworker

import (
	"bytes"
	"testing"
)

func TestWriteFile(t *testing.T) {
	tests := []struct {
		name  string
		data  []interface{}
		wantW string
	}{
		{
			name:  "String",
			data:  []interface{}{"string"},
			wantW: "string",
		},
		{
			name:  "Numbers",
			data:  []interface{}{1},
			wantW: "",
		},
		{
			name:  "Few strings",
			data:  []interface{}{"string1", "string2"},
			wantW: "string1string2",
		},
		{
			name:  "Strings and numbers",
			data:  []interface{}{"string", 1, 2, " ", "number"},
			wantW: "string number",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			WriteFile(w, tt.data...)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("WriteFile() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
