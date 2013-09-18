gopwgen
=======

A simple password generator written in Go

Usage
-----
```
Usage: gopwgen <OPTION> [length] [number] (length and number optional)

OPTIONS (MUST SPECIFY ONE!)  
s | symbols      Alphanumeric + symbols (NOT FOR MYSQL!)  
a | alpha        Alphanumeric only  
p | phpmyadmin   Generate phpMyAdmin Blowfish secret (for cookie auth)  
w | wordpress    Generate Wordpress encryption keys (wp-config.php)  
h | help         Display this usage information  

If no length or number are defined (and you haven't changed the code), a default length of 19 and number of 1 will be used.
```
Usage can also be shown by typing `gopwgen h` or `gopwgen help`
