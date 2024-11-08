package urand

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func RandomSixDigit() string {
	return RandomString(6)
}

func RandomTenDigit() string {
	return RandomString(10)
}

func RandomTwelveDigit() string {
	return RandomString(12)
}

func RandomSixDigitOTP() int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	opt := 100000 + rng.Intn(900000)
	return opt
}
