package tools

import (
	// "encoding/json"
	"fmt"
	// "io"
	// "log"
	// "os"
	// "testing"
	"math/rand"
	"time"
)

type envs struct {
	PINGONE_USERNAME string `json:"PINGONE_USERNAME"`
	PINGONE_PASSWORD string `json:"PINGONE_PASSWORD"`
}

func PrintHeader(str string) {
	fmt.Printf("***START:%v***\n", str)
}

func PrintFooter(str string) {
	fmt.Printf("***END:%v***\n", str)
}

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString(length int) string {
	return StringWithCharset(length, charset)
}
