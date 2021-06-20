package main

import (
	"bufio"
	"fmt"
	"os"
)

var defaultAlphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 !?,."

func main() {
	// check if alphabet was overridden
	alphabet := defaultAlphabet
	if env, ok := os.LookupEnv("VIGENERE_ALPHABET"); ok {
		alphabet = env
	}

	if len(os.Args) < 2 {
		fmt.Println("Encryption key is required!")
		os.Exit(1)
	}

	key := os.Args[1]

	// read input
	var text string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text += scanner.Text()
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}

	cipher := NewVigenereCipher(key, alphabet)
	encr, err := cipher.Encrypt([]rune(text))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(encr))
}
