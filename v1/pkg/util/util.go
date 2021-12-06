package util

import (
	"fmt"
	"math/rand"
	"time"
)

// EnvType donde se esta desplegando la app
type EnvType string

const (
	EnvProduction EnvType = "production"
	EnvDevelop    EnvType = "develop"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandString retorna un string oleatorio
func RandString() string {
	now := time.Now().UTC().Unix()
	b := make([]byte, 7)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return fmt.Sprintf("%s%d", string(b), now)
}

// RandStringn retorna un string oleatorio de len=strLen
func RandStringn(strLen int) string {
	b := make([]byte, strLen)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// GetLocationFromIP obtiene una localizacion a partir del ip
func GetLocationFromIP(ip string) string {
	return "country-city-town"
}
