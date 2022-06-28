package randName

import (
	"math/rand"
	"strings"
	"time"
)

func New() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyz")
	length := rand.Intn(7) + 5
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	str = strings.ToUpper(str[0:1]) + str[1:]
	return str
}
