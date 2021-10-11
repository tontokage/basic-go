package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	mode := flag.Int("mode", 256, "256 or 384 or 512")
	flag.Parse()

	fmt.Print("input wait\n")

	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		switch *mode {
		case 256:
			sum256 := sha256.Sum256(input.Bytes())
			fmt.Printf("%x\n", sum256)
		case 384:
			sum384 := sha512.Sum384(input.Bytes())
			fmt.Printf("%x\n", sum384)
		case 512:
			sum512 := sha512.Sum512(input.Bytes())
			fmt.Printf("%x\n", sum512)
		default:
			fmt.Fprintf(os.Stderr, "invalid mode: %d\n", *mode)
			os.Exit(1)
		}
	}

}
