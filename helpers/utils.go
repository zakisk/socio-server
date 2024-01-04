package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// Generate a salt with cost 14 (you can adjust the cost according to your security requirements)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
