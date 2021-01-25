package main

import (
	"angopipe"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	gcm, err := angopipe.Prepare()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	source, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Encryption Error: %v\n", err)
		os.Exit(1)
	}

	nonce := make([]byte, 12)
	io.ReadFull(rand.Reader, nonce)
	os.Stdout.Write(nonce)
	result := gcm.Seal(nil, nonce, source, nil)
	os.Stdout.Write(result)
}
