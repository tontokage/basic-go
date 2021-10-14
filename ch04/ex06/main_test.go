package main

import (
	"reflect"
	"testing"
)

func TestIsASCII(t *testing.T) {
	tests := []struct {
		name string
		b    uint8
		want bool
	}{
		{name: "ascii", b: []byte("a")[0], want: true},
		{name: "ascii", b: []byte("a")[0], want: true},
		{name: "2byte", b: []byte("")[0], want: false},
		{name: "3byte", b: []byte("あ")[0], want: false},
		{name: "4byte", b: []byte("𐅃")[0], want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := isASCII(test.b); got != test.want {
				t.Errorf("isASCII = '%v', want = '%v'", got, test.want)
			}
		})
	}

}

func TestIs2Byte(t *testing.T) {
	tests := []struct {
		name string
		b    uint8
		want bool
	}{
		{name: "ascii", b: []byte("a")[0], want: false},
		{name: "2byte", b: []byte("")[0], want: true},
		{name: "3byte", b: []byte("あ")[0], want: false},
		{name: "4byte", b: []byte("𐅃")[0], want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := is2Byte(test.b); got != test.want {
				t.Errorf("is2Byte = '%v', want = '%v'", got, test.want)
			}
		})
	}

}

func TestIs3Byte(t *testing.T) {
	tests := []struct {
		name string
		b    uint8
		want bool
	}{
		{name: "ascii", b: []byte("a")[0], want: false},
		{name: "2byte", b: []byte("")[0], want: false},
		{name: "3byte", b: []byte("あ")[0], want: true},
		{name: "4byte", b: []byte("𐅃")[0], want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := is3Byte(test.b); got != test.want {
				t.Errorf("is3Byte = '%v', want = '%v'", got, test.want)
			}
		})
	}

}

func TestIs4Byte(t *testing.T) {
	tests := []struct {
		name string
		b    uint8
		want bool
	}{
		{name: "ascii", b: []byte("a")[0], want: false},
		{name: "2byte", b: []byte("")[0], want: false},
		{name: "3byte", b: []byte("あ")[0], want: false},
		{name: "4byte", b: []byte("𐅃")[0], want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := is4Byte(test.b); got != test.want {
				t.Errorf("is4Byte = '%v', want = '%v'", got, test.want)
			}
		})
	}

}
func TestCompressSpaces(t *testing.T) {
	type want struct {
		b   []byte
		err error
	}
	tests := []struct {
		name string
		args []byte
		want want
	}{
		{
			name: "スペースありだが隣接しない",
			args: []byte("良い\t天気ですね"),
			want: want{b: []byte("良い\t天気ですね"), err: nil},
		},
		{
			name: "スペースありで隣接する",
			args: []byte("良い\t\t天気ですね"),
			want: want{b: []byte("良い 天気ですね"), err: nil},
		},
		{
			name: "スペースありで隣接する 最初の文字",
			args: []byte("\t\t良い天気ですね"),
			want: want{b: []byte(" 良い天気ですね"), err: nil},
		},
		{
			name: "スペースありで隣接しない 最後の文字",
			args: []byte("良い天気ですね\t"),
			want: want{b: []byte("良い天気ですね"), err: nil},
		},
		{
			name: "スペースありで隣接する 最後の文字",
			args: []byte("良い天気ですね\t\t"),
			want: want{b: []byte("良い天気ですね "), err: nil},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := compressSpaces(test.args)
			if !reflect.DeepEqual(got, test.want.b) || err != test.want.err {
				t.Errorf("compressSpaces = '%s', want = '%s'", got, test.want.b)
			}
		})
	}

}
