package service

import (
	"crypto/rand"
	"fmt"

	domainio "github.com/Avon11/Chotu-Go/internal/DomainIo"
)

func CreateShortUrl(oldUrl string) (shortCode *domainio.ShortCodeDomain, errResp *domainio.ErrorResponse) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomBytes := make([]byte, 6)
	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 6; i++ {
		randomBytes[i] = charset[int(randomBytes[i])%len(charset)]
	}
	fmt.Println("Random string : ", string(randomBytes))
	return
}
