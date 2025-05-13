package utils

import (
	"log"
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

func GenerateRequestID(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func GetRequestID(ctx *fiber.Ctx) string {
	requestID := ctx.Locals("requestID")
	if requestID != nil {
		if id, ok := requestID.(string); ok {
			return id
		}
	}
	return "unknown"
}

func LogRegister(ctx *fiber.Ctx, component string, message string) {
	log.Printf("[RequestID: %s] %s: ", GetRequestID(ctx), component+" "+message)
}

func LogErrorRegister(ctx *fiber.Ctx, component string, err error) {
	log.Printf("[ERROR] [RequestID: %s] %s: %v", GetRequestID(ctx), component, err)
}
