package utils

import (
	"os"
	"regexp"
	"strconv"

	"github.com/dongri/phonenumber"
)

// GetEnv gets the environment variable
func GetEnv(key string) string {
	return os.Getenv(key)
}

// GetEnvWithDefault gets the environment variable with a default value
func GetEnvWithDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// Convert String to uint
func StringToUint(s string) (uint, error) {
	res, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(res), nil
}

// Phone number validation
func ValidatePhoneNumber(phoneNumber string) string {
	re := regexp.MustCompile(`^(\+\d{1,3}[- ]?)?\d{10,11}$`)
	if re.MatchString(phoneNumber) {
		phone := phonenumber.Parse(phoneNumber, "ID")
		return phone
	}
	return ""
}
