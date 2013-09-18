// Sal's Random Password Generator
// --------------------------------
// My very first Go application.
// Usage is printed by running 'salpwgen h'

package main

import (
    "fmt"
    "math/rand"
    "time"
    "os"
    "flag"
    "strconv"
)

// Function to print usage information
func usage() {
    fmt.Println("Sal's Random Password Generator")
    fmt.Println("-------------------------------")
    fmt.Println("Usage: salpwgen <OPTION> [length] [number] (length and number optional)\n")
    fmt.Println("OPTIONS (MUST SPECIFY ONE!):")
    fmt.Println("s | symbols      Alphanumeric + symbols (NOT FOR MYSQL!)")
    fmt.Println("a | alpha        Alphanumeric only")
    fmt.Println("p | phpmyadmin   Generate phpMyAdmin Blowfish secret (for cookie auth)")
    fmt.Println("w | wordpress    Generate Wordpress encryption keys (wp-config.php)")
    fmt.Println("h | help         Display this usage information\n")
    fmt.Println("If no length or number are defined (and you haven't changed the code), a default length of 19 and number of 1 will be used.")
    os.Exit(0)
}

func main() {
    // Alphanumeric values
    alphanumeric := "abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"

    // Symbols + alpha values
    symbols := "abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789-_!@#$%^&*/\\()_+{}|:<>?="

    // Set default length+number if no args, or
    // just read the args and populate one/both
    var length int
    var number int
    
    flag.Parse()
    if flag.Arg(1) == "" && flag.Arg(2) == "" {
        length = 19
        number = 1
    } else if flag.Arg(2) == "" {
        s1 := flag.Arg(1)
        value1, err1 := strconv.Atoi(s1)
        if err1 != nil {
            fmt.Println(err1)
            os.Exit(2)
        }
        length = value1
        number = 1
    } else {
        s1 := flag.Arg(1)
        s2 := flag.Arg(2)
        value1, err1 := strconv.Atoi(s1)
        if err1 != nil {
            fmt.Println(err1)
            os.Exit(2)
        }
        value2, err2 := strconv.Atoi(s2)
        if err2 != nil {
            fmt.Println(err2)
            os.Exit(2)
        }
        length = value1
        number = value2
    }

    // Seed the RNG
    rand.Seed(time.Now().UnixNano())

    // Create the password string
    password := make([]byte, length)

    // Generate alphanumeric password(s)
    if flag.Arg(0) == "a" || flag.Arg(0) == "alpha" {
        for i := 0; i < number; i++ {
            for j := 0; j < length; j++ {
                password[j] = alphanumeric[rand.Intn(len(alphanumeric))]
            }

            fmt.Println(string(password))
        }
    }

    // Generate alphanumeric + symbols password(s)
    if flag.Arg(0) == "s" || flag.Arg(0) == "symbols" { 
        for i := 0; i < number; i++ {
            for j := 0; j < length; j++ {
                password[j] = symbols[rand.Intn(len(symbols))]
            }   
            
            fmt.Println(string(password))
        }
    }

    // Generate phpMyAdmin blowfish secret
    if flag.Arg(0) == "p" || flag.Arg(0) == "phpmyadmin" {
        password := make([]byte, 64)
        for j := 0; j < 64; j++ {
            password[j] = symbols[rand.Intn(len(symbols))]
        }

        fmt.Println(string(password))
        
    }

    // Generate WordPress encryption secrets
    if flag.Arg(0) == "w" || flag.Arg(0) == "wordpress" {
        password1 := make([]byte, 64)
        password2 := make([]byte, 64)
        password3 := make([]byte, 64)
        password4 := make([]byte, 64)
        password5 := make([]byte, 64)
        password6 := make([]byte, 64)
        password7 := make([]byte, 64)
        password8 := make([]byte, 64)

        for j := 0; j < 64; j++ {
            password1[j] = symbols[rand.Intn(len(symbols))]
        }
        for j := 0; j < 64; j++ {
            password2[j] = symbols[rand.Intn(len(symbols))]
        }
        for j := 0; j < 64; j++ {
            password3[j] = symbols[rand.Intn(len(symbols))]
        }
        for j := 0; j < 64; j++ {
            password4[j] = symbols[rand.Intn(len(symbols))]
        }
        for j := 0; j < 64; j++ {
            password5[j] = symbols[rand.Intn(len(symbols))]
        }
        for j := 0; j < 64; j++ {
            password6[j] = symbols[rand.Intn(len(symbols))]
        }
        for j := 0; j < 64; j++ {
            password7[j] = symbols[rand.Intn(len(symbols))]
        }
        for j := 0; j < 64; j++ {
            password8[j] = symbols[rand.Intn(len(symbols))]
        }

        fmt.Printf("define('AUTH_KEY',\t\t'%s');\n",password1)
        fmt.Printf("define('SECURE_AUTH_KEY',\t'%s');\n",password2)
        fmt.Printf("define('LOGGED_IN_KEY',\t\t'%s');\n",password3)
        fmt.Printf("define('NONCE_KEY',\t\t'%s');\n",password4)
        fmt.Printf("define('AUTH_SALT',\t\t'%s');\n",password5)
        fmt.Printf("define('SECURE_AUTH_SALT',\t'%s');\n",password6)
        fmt.Printf("define('LOGGED_IN_SALT',\t'%s');\n",password7)
        fmt.Printf("define('NONCE_SALT',\t\t'%s');\n",password8)

    }

    // Print usage info
    if flag.Arg(0) == "h" || flag.Arg(0) == "help" || flag.Arg(0) == "" {
        usage()
    }

}
