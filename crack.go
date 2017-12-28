// apparently I always need this
package main

// import all the libraries I'll need
import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

// define a string variable "cipher" from the first command line arg
var cipher string = os.Args[1]

func main() {
	// give a heads up that we're starting
	fmt.Println("[+] Cracking...")

	c := make(chan string)
	// C-style for loop
	for i := 1; i < 99999999; i++ {
		// use strconv to convert integer to string,
		// hash the string, then compare it to our provided hash
		go GetMD5Hash(i, cipher, c)
	}

	i := <-c
	fmt.Println("[!] The code has been cracked: ", i)
	os.Exit(0)
}

// function that takes a string and returns the md5 of that string
func GetMD5Hash(i int, cipher string, c chan string) {
	text := strconv.Itoa(i)
	hasher := md5.New()
	hasher.Write([]byte(text))
	res := hex.EncodeToString(hasher.Sum(nil))
	if res == cipher {
		c <- res
	}
}
