package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	if len(os.Args) == 1 {
		help := "The program generates a time-based one-time password based in the given secret and current time\n"
		help += "ERROR: The program expects a secret as command-line argument.\n"
		help += "The secret must be a base32-encoded string."
		fmt.Fprintln(os.Stderr, help)
		os.Exit(1)
	}

	totp, err := GetTOTP(os.Args[1], time.Now().Unix())
	if err == nil {
		fmt.Print(totp)
	} else {
		fmt.Fprint(os.Stderr, err)
	}
}
