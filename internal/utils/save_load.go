package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const (
	Dir = "blockChainData/"
)

func Save(data []byte, fileName string) error {
	os.Mkdir(Dir, 0755)
	file, err := os.OpenFile(
		fmt.Sprintf("%s/%s.json", Dir, fileName),
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
		fmt.Sprintf("%s/%s.json", Dir, fileName),
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
