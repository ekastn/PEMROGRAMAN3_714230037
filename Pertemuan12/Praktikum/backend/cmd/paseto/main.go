package main

import (
	"aidanwoods.dev/go-paseto"
	"fmt"
)

func main() {
	privateKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := privateKey.Public()

	fmt.Println("Private Key (hex):", privateKey.ExportHex())
	fmt.Println("Public Key (hex):", publicKey.ExportHex())
}
