package service

import (
	"crypto/rand"
	"log"
	"time"

	db "github.com/Avon11/ShrinkRay/internal/DB"
	domainio "github.com/Avon11/ShrinkRay/internal/DomainIo"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

const prefix = "https://ShrinkRay.com/"

type ShortCodeService struct {
	RedisClient *redis.Client
}

func NewCodeService(redisClient *redis.Client) *ShortCodeService {
	return &ShortCodeService{
		RedisClient: redisClient,
	}
}

func (s *ShortCodeService) CreateShortUrl(c *gin.Context, oldUrl string) (shortCode *domainio.ShortCodeDomain, errResp *domainio.ErrorResponse) {
	shortUrlCode := ""
	for {
		shortUrlCode = CreateShortCode()
		// checking for duplicate url
		notFound := s.ValidateDupliateKey(c, shortUrlCode)
		if !notFound {
			break
		}
	}

	// saving in DB
	err := AddShortCodeService(oldUrl, shortUrlCode)
	if err != nil {
		errResp = &domainio.ErrorResponse{
			Error:   err,
			ErrMsg:  "internal server error",
			ErrCode: 500,
		}
		log.Println("Error while creating short code", err)
		return
	}

	// saving in cache
	err = s.SaveInCache(c, shortUrlCode, oldUrl)
	if err != nil {
		errResp = &domainio.ErrorResponse{
			Error:   err,
			ErrMsg:  "internal server error",
			ErrCode: 500,
		}
		log.Println("Error while creating short code", err)
		return
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
		log.Println("Error while reading random bytes", err)
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
		log.Println("Error while storing short code to DB", err)
	}
	return
}

func (s *ShortCodeService) ValidateDupliateKey(c *gin.Context, shortCode string) (exist bool) {
	code, err := s.GetFromCache(c, shortCode)
	if err != nil && err != redis.Nil {
		log.Println("Error while validating shortcode", err)
	}
	if len(code) > 0 {
		return true
	}
	exist, err = db.CheckForDuplicateKey(shortCode)
	if err != nil {
		log.Println("Error while validating shortcode", err)
	}
	return
}

func (s *ShortCodeService) RedirectUrl(c *gin.Context, shortUrl string) (redirectUrl domainio.RedirectUrl, errResp *domainio.ErrorResponse) {

	redirectUrl.Code = 200

	if len(shortUrl) < 6 {
		log.Printf("Invalid shortCode: too short. Length: %d", len(shortUrl))
		errResp = &domainio.ErrorResponse{
			ErrCode: 400,
			ErrMsg:  "Invalid shortCode: too short",
		}
		return
	}

	shortCode := shortUrl[:6]
	redirectUrl.ShortCode = shortCode

	if len(shortUrl) > 6 && shortUrl[len(shortUrl)-1] == '_' {
		redirectUrl.Code = 201
		log.Printf("Special case detected. Setting redirect code to 201")
	}

	cacheKey := shortCode
	cachedResponse, err := s.GetFromCache(c, cacheKey)
	if err != nil {
		if err == redis.Nil {
			log.Printf("Cache miss for key: %s", cacheKey)
		} else {
			log.Printf("Error fetching Url from Redis: %v", err)
			errResp = &domainio.ErrorResponse{
				ErrCode: 500,
				ErrMsg:  "Internal server error",
			}
			return
		}
	}

	if len(cachedResponse) > 0 {
		log.Printf("Cache hit. Returning cached URL for key: %s", cacheKey)
		redirectUrl.Url = cachedResponse
		return
	}

	url, err := db.GetUrlByShortCode(shortCode)
	log.Printf("Err e: %s", err)
	if err != nil {
		log.Printf("Error fetching Url from DB: %v", err)
		errResp = &domainio.ErrorResponse{
			ErrCode: 404,
			ErrMsg:  "Url not found",
		}
		return
	}

	err = s.SaveInCache(c, cacheKey, url)
	if err != nil {
		log.Printf("Error saving URL to cache: %v", err)
	}

	redirectUrl.Url = url
	return
}

func (s *ShortCodeService) SaveInCache(c *gin.Context, cacheKey string, url string) error {
	return s.RedisClient.Set(c, cacheKey, url, 24*7*time.Hour).Err()
}

func (s *ShortCodeService) GetFromCache(c *gin.Context, cacheKey string) (url string, err error) {
	url, err = s.RedisClient.Get(c, cacheKey).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return
}
