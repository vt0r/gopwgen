# gopwgen

A simple password generator written in Go

## Usage

``` txt
Usage: gopwgen [OPTION] [length | num dice words] [num pwds]

Options:
  -H    Hexadecimal only (abcdef0123456789)
  -a    Alphanumeric only
  -d    Diceware passphrase (Choose random words from a list)
  -p    Generate phpMyAdmin Blowfish secret (for cookie auth)
  -s    Alphanumeric + symbols (NOT FOR MYSQL!)
  -v    Print the version number and exit
  -w    Generate WordPress encryption salts for use in wp-config.php

If no password length or number are defined (and you haven't changed the code),
a default length of 19 and number of 1 will be used.

For Diceware, the default number of words per passphrase is 6.
```

## Performance

You'll generally get the best performance by compiling this script into a binary first, which is as simple as running: `go build gopwgen.go`  
You can then run `./gopwgen` or simply `gopwgen` if you copy the binary to `~/bin` or `/usr/bin`.

## Thanks

Special thanks goes to:  
  
[Jacques Fuentes](https://github.com/jpfuentes2)  
[Nikhil Narula](https://github.com/nn2242)  
[Bodie Solomon](https://github.com/binary132)  
  
for pointing out spots where I'd goofed or where code could be shrunk/improved.
