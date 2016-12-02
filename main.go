package main

import "crypto/aes"
import "fmt"
import "os"

// AES-128 key
var key = []byte{
	0, 0, 0, 0,
	0, 0, 0, 0,
	0, 0, 0, 0,
	0, 0, 0, 0,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Not enough arguments\n")
		return
	}

	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// Read first argument from the command line and pad with zeros
	// to a multiple of the block size if necessary.
	blkSize := cipher.BlockSize()
	buf := []byte(os.Args[1])
	if len(buf)%blkSize != 0 {
		buf = append(buf, make([]byte, blkSize-len(buf)%blkSize)...)
	}

	// Encrypt and decrypt in place
	fmt.Printf("Plaintext: %s\n", buf)
	fmt.Printf("      Raw: %v\n\n", buf)
	cipher.Encrypt(buf, buf)
	fmt.Printf("Encrypted: %v\n\n", buf)
	cipher.Decrypt(buf, buf)
	fmt.Printf("Decrypted: %s\n", buf)
	fmt.Printf("      Raw: %v\n", buf)
}
