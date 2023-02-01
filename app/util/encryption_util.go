package util

import "golang.org/x/crypto/bcrypt"

// GenerateHashFromPassword generates password hash
func GenerateHashFromPassword(d string) string {
	h, err := bcrypt.GenerateFromPassword([]byte(d), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(h)
}

// CompareHashAndPassword compares pass with hash
func CompareHashAndPassword(h string, d string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(d))
	if err != nil {
		return false
	}
	return true
}
