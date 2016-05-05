gopwgen
=======

A simple password generator written in Go

Usage
-----
```
Usage: gopwgen <OPTION> [length] [number] (length and number optional)

OPTIONS (MUST SPECIFY ONE!)
-s               Alphanumeric + symbols (NOT FOR MYSQL!)
-a               Alphanumeric only
-H               Hexadecimal only (abcdef0123456789)
-p               Generate phpMyAdmin Blowfish secret (for cookie auth)
-w               Generate WordPress encryption keys (wp-config.php)
-h               Display this usage information

If no length or number are defined (and you haven't changed the code),
a default length of 19 and number of 1 will be used.
```

Performance
-----------
You'll generally get the best performance by compiling this script into a binary first, which is as simple as running: `go build gopwgen.go`  
You can then run `./gopwgen` or simply `gopwgen` if you copy the binary to `~/bin` or `/usr/bin`.


Thanks
------
Special thanks goes to:  
  
[Jacques Fuentes](https://github.com/jpfuentes2)  
[Nikhil Narula](https://github.com/nn2242)  
[Bodie Solomon](https://github.com/binary132)  
  
for pointing out spots where I'd goofed or where code could be shrunk/improved.
