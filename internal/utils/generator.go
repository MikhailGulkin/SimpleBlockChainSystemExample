package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GenerateId() string {
	var result strings.Builder
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 20; i++ {
		if i%4 == 0 && i != 0 {
			result.WriteString("-")
		}
		if i%2 == 0 {
			result.WriteString(strconv.Itoa(random.Intn(10)))
		} else {
			result.WriteString(string(rune(random.Intn(26) + 65)))
		}
	}
	return result.String()
}
