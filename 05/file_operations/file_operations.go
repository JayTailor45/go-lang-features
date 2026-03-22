package file_operations

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteBalanceToFile(value float64, fileName string) {
	os.WriteFile(fileName, []byte(fmt.Sprint(value)), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("Failed to find file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse storaged value")
	}
	return value, nil
}
