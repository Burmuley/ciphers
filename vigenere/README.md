# Vigenere encryption tool

This tool is just a simple dumb implementation of the [Vigenere cipher](https://en.wikipedia.org/wiki/Vigenère_cipher).

The default alphabet includes lower and upper case latin letters, numbers and some special characters:
```
abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 !?,.
```

You can override this alphabet by setting a value to environment variable `VIGENERE_ALPHABET`.

Letters not present in the defined alphabet will be silently dropped.

Use is simple:
```bash
$ go build -o vigenere
$ ./vigenere <encryption key>
```
Where `encryption key` is a key word used to set the mapping on the encryption table where each row is shifted according
to [Ceasar cipher](https://en.wikipedia.org/wiki/Caesar_cipher)

Command above will expect an input on `stdin` (you can end your input by pressing Control+D (CTRL+D for obsolete ugly operating systems)).

You can also pipe a file contents:
```bash
$ ./vigenere <encryption key> < plaintext.txt
```

As the algorithm implements simple reversible approach, to decrypt a message simply run the tool with the same key and encrypted text on input. 
