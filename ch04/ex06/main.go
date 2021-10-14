package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("良い  天気ですね")
	ret, _ := compressSpaces(b)
	fmt.Printf("%s\n", b)
	fmt.Printf("%s\n", ret)
}

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

func compressSpaces(b []byte) ([]byte, error) {
	i, j := 0, 0
	var previous []byte

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

		if i == 0 { // init
			if isSpace(b[i : i+size]) {
				previous = []byte(" ")
				b[i] = ' '
				j += 1
			} else {
				previous = b[i : i+size]
				j += size
			}
			i += size
			continue
		}

		currentIsSpace := isSpace(b[i : i+size])
		if !(isSpace(previous) && currentIsSpace) {
			if currentIsSpace {
				b[j] = ' '
				j += 1
			} else {
				copy(b[j:j+size], b[i:i+size])
				j += size
			}
		}

		previous = b[i : i+size]
		i += size
	}

	return b[:j], nil
}

func isSpace(b []byte) bool {
	r, _ := utf8.DecodeRune(b)
	switch uint32(r) {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	}
	return false
}
