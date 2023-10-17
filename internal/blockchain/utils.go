package blockchain

import (
	"encoding/json"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func isBalanceSufficient(senderBalance int64, amount int64) bool {
	return senderBalance >= amount
}
func Save(data []byte, fileName string) error {
	os.Mkdir(Dir, 0755)
	file, err := os.OpenFile(
		fileName,
		os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644,
	)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func Load(object interface{}, fileName string) error {
	file, err := os.OpenFile(
		fileName,
		os.O_RDONLY, 0644,
	)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	byteValue, err := io.ReadAll(file)

	if err != nil {
		return err
	}
	return json.Unmarshal(byteValue, object)
}
func GenerateTransactionId() string {
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
