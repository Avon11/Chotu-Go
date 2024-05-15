package service

import (
	"crypto/rand"
	"log"

	db "github.com/Avon11/Chotu-Go/internal/DB"
	domainio "github.com/Avon11/Chotu-Go/internal/DomainIo"
)

const prefix = "https://chotu.com/" // Replace with frontend url ---- TODO

func CreateShortUrl(oldUrl string) (shortCode *domainio.ShortCodeDomain, errResp *domainio.ErrorResponse) {
	shortUrlCode := ""
	for {
		shortUrlCode = CreateShortCode()
		notFound := ValidateDupliateKey(shortUrlCode)
		if !notFound {
			break
		}
	}
	err := AddShortCodeService(oldUrl, shortUrlCode)
	if err != nil {
		errResp = &domainio.ErrorResponse{
			Error:   err,
			ErrMsg:  "internal server error",
			ErrCode: 500,
		}
		log.Fatalln("Error while creating short code", err)
	}
	shortUrl := prefix + shortUrlCode
	shortCode = &domainio.ShortCodeDomain{
		Url: shortUrl,
	}
	return
}

func CreateShortCode() (code string) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomBytes := make([]byte, 6)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatalln("Error while reading random bytes", err)
	}
	for i := 0; i < 6; i++ {
		randomBytes[i] = charset[int(randomBytes[i])%len(charset)]
	}
	code = string(randomBytes)
	return
}

func AddShortCodeService(url, shortCode string) (err error) {
	err = db.AddShortCode(shortCode, url)
	if err != nil {
		log.Fatalln("Error while storing short code to DB", err)
	}
	return
}

func ValidateDupliateKey(shortCode string) (exist bool) {
	exist, err := db.CheckForDuplicateKey(shortCode)
	if err != nil {
		log.Fatalln("Error while validating shortcode", err)
	}
	return
}
