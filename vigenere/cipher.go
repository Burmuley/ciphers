package main

import (
	"fmt"
	"strings"
)

type VigenereCipher struct {
	alphabet []rune
	cipher   map[rune][]rune
	key      []rune
	index    map[rune]int
}

func (c VigenereCipher) Encrypt(input []rune) ([]rune, error) {
	keyLen := len(c.key)
	out := make([]rune, len(input))
	input = filterInput(input, c.alphabet)

	for i, r := range input {
		curKey := c.key[i%keyLen]
		keyIndex, ok := c.index[curKey]

		if !ok {
			return nil, fmt.Errorf("key index %d not found", curKey)
		}

		out[i] = c.cipher[r][keyIndex]
	}

	return out, nil
}

func (c VigenereCipher) Decrypt(input []rune) ([]rune, error) {
	return c.Encrypt(input)
}

func NewVigenereCipher(key string, alphabet string) *VigenereCipher {
	c := &VigenereCipher{
		key:      []rune(strings.ToLower(key)),
		alphabet: []rune(alphabet),
	}

	c.cipher = makeCipher(alphabet)
	updateIndex(c)
	return c
}

func updateIndex(c *VigenereCipher) {
	c.index = make(map[rune]int)
	for i, l := range c.alphabet {
		c.index[l] = i
	}
}

func makeCipher(alphabet string) map[rune][]rune {
	letters := []rune(alphabet)
	cipher := make(map[rune][]rune)

	for _, l := range letters {
		letters = shiftSlice(letters)
		cipher[l] = letters
	}

	return cipher
}

func shiftSlice(s []rune) []rune {
	ns := make([]rune, len(s))
	ns[0] = s[len(ns)-1]

	for i := 1; i < len(ns); i++ {
		ns[i] = s[i-1]
	}

	return ns
}

func isInList(r rune, l []rune) bool {
	for _, i := range l {
		if r == i {
			return true
		}
	}
	return false
}

func filterInput(input []rune, alphabet []rune) []rune {
	out := make([]rune, 0, len(input))
	for _, i := range input {
		if isIn := isInList(i, alphabet); isIn {
			out = append(out, i)
		}
	}

	return out
}
