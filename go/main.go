package main

import (
	"encoding/base64"
	"fmt"

	"github.com/donsdai/utils/utils"
)

func main() {
	saltbyte, _ := utils.GenerateRandomBytes(16)
	b64Salt := base64.RawStdEncoding.EncodeToString(saltbyte)
	fmt.Println("Salt for hashing: " + b64Salt)

	bcrypt, _ := utils.HashPassword(utils.Cbcrypt, "12345678", b64Salt)
	fmt.Println("Hash password with bcrypt: " + bcrypt)

	argon, _ := utils.HashPassword(utils.Cargon2, "12345678", b64Salt)
	fmt.Println("Hash password with argon2: " + argon)

	uuid, _ := utils.GenerateUUID()
	fmt.Println("Generating UUID: " + uuid)
}
