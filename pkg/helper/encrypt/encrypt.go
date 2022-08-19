package encrypt

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func CheckPassword(password string, hashedPassword string) bool {
	result := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return result == nil
}
