package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadFromJsonFile[T any](filepath string) (T, error) {
	var data T

	jsonFile, err := os.Open(filepath)

	if err != nil {
		return data, fmt.Errorf("failed to open file : %v", err)
	}

	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)

	err = decoder.Decode(data)

	if err != nil {
		return data, fmt.Errorf("failed to decode json : %v", err)
	}

	return data, nil
}
