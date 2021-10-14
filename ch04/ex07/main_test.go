package main

import (
	"reflect"
	"testing"
)

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func TestCompressSpaces(t *testing.T) {
	tests := []struct {
		name string
		args []byte
		want []byte
	}{
		{
			name: "1byte",
			args: []byte("aiu"),
			want: []byte("uia"),
		},
		{
			name: "2byte",
			args: []byte("あいう"),
			want: []byte("ういあ"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, _ := reverseUTF8(test.args)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("reverseUTF8 = %s, want = %s", got, test.want)
			}
		})
	}

}
