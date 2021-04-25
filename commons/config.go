package commons

import (
	"log"
	"os"
	"strconv"
)

func GetInt(name string, defaultVal int) int {
	s := os.Getenv(name)
	if len(s) == 0 {
		return defaultVal
	}

	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("env %s must be an interger", name)
	}
	return num
}
