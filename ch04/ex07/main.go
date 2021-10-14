package main

import (
	"errors"
	"fmt"
)

func isASCII(b uint8) bool {
	return b&(1<<7) == 0
}

func is2Byte(b uint8) bool {
	return !isASCII(b) && b&(1<<6) != 0 && b&(1<<5) == 0
}

func is3Byte(b uint8) bool {
	return !isASCII(b) && b&(1<<6) != 0 && b&(1<<5) != 0 && b&(1<<4) == 0
}

func is4Byte(b uint8) bool {
	return !isASCII(b) && b&(1<<6) != 0 && b&(1<<5) != 0 && b&(1<<4) != 0 && b&(1<<3) == 0
}

func main() {
	b := []byte("あいう")
	a, _ := reverseUTF8(b)
	fmt.Printf("%s\n", a)

}

func reverseUTF8(b []byte) ([]byte, error) {
	i := 0
	for i < len(b) {
		var size int
		switch {
		case isASCII(b[i]):
			size = 1
		case is2Byte(b[i]):
			size = 2
		case is3Byte(b[i]):
			size = 3
		case is4Byte(b[i]):
			size = 4
		default:
			return nil, errors.New(fmt.Sprintf("invalid code: %x %[1]b\n", b[i]))
		}

		// 文字単位でreverseする
		// 3byteなら 0,1,2 -> 2,1,0
		// 最後に全部reverseするので元の並びに戻る
		if size > 1 {
			reverse(b[i : i+size])
		}

		i += size
	}

	// 全体をreverseする
	reverse(b)
	return b, nil
}

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
