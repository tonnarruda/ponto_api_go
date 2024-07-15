package helper

import (
	"fmt"
	"log"
	"strconv"
)

func GenerateNextCode(lastCode string) string {
	if lastCode == "" {
		return "0001"
	}

	lastNum, err := strconv.Atoi(lastCode)
	if err != nil {
		log.Printf("Error converting last code to integer: %v", err)
		return ""
	}

	nextNum := lastNum + 1
	nextCode := fmt.Sprintf("%04d", nextNum)

	return nextCode
}
