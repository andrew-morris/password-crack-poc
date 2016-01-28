// apparently I always need this
package main

// import all the libraries I'll need
import ( 					
	"fmt"
    "os"
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

// define a string variable "cipher" from the first command line arg
var cipher string = os.Args[1] 		

func main() {
	// give a heads up that we're starting
	fmt.Println("[+] Cracking...")
    // C-style for loop
	for i := 1; i < 99999999; i++ {
    	// use strconv to convert integer to string, 
        // hash the string, then compare it to our provided hash
		if GetMD5Hash(strconv.Itoa(i)) == cipher { 
        	// if we get a match, print it out and exit with a zero exit code
			fmt.Printf("[!] The code has been cracked: %d\n", i)
			os.Exit(0)
		}
	}
}

// function that takes a string and returns the md5 of that string
func GetMD5Hash(text string) string { 			
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}