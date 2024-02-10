package util

import (
	"crypto/rand"
	"math/big"
	"strings"
	"time"
)

type HelperConfig struct{}

func (h HelperConfig) OnRandomChar(countChar int) string {
	characters := "ABCDEFGHIJKLMNPQRSTUVWXYZ123456789#@!*"
	result := make([]byte, countChar)
	for i := 0; i < countChar; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		result[i] = characters[index.Int64()]
	}
	return string(result)
}

func (h HelperConfig) OnRandomNumber(countChar int) string {
	characters := "0123456789"
	result := make([]byte, countChar)
	for i := 0; i < countChar; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		result[i] = characters[index.Int64()]
	}
	return string(result)
}

func (h HelperConfig) OnGeneratePassword(countChar int) string {
	var RAND bool
	RAND = false
	var GENERATE_PASS string
	for !RAND {
		GENERATE_PASS = h.OnRandomChar(countChar)
		if strings.ContainsAny(GENERATE_PASS, "@!#") {
			RAND = true
		}
	}
	return GENERATE_PASS
}

func (h HelperConfig) OnCurrentTime() int64 {
	return time.Now().Unix()
}

func (h HelperConfig) OnGetDateNow() string {
	now := time.Now()
	return now.Format("2006-01-02")
}

func (h HelperConfig) OnGetDateTimeNow() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}
