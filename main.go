package main

import (
	"fmt"
	"time"

	notation "github.com/notaryproject/notation-go"
)

func main() {
	n := notation.SignOptions{
		SignatureMediaType: "application/cose",
		ExpiryDuration:     10 * time.Hour,
	}
	fmt.Printf("n: %v\n", n)
}
