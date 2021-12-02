package auth

import "golang.org/x/crypto/bcrypt"

// GeneratePassword retorna una nueva clave
func GeneratePassword(pwd []byte) (string, error) {
	hash, e := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if e != nil {
		return "", e
	}
	return string(hash), nil
}

// ComparePasswords compara passwords
func ComparePasswords(hashedPwd string, plainPwd []byte) error {
	byteHash := []byte(hashedPwd)
	e := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if e != nil {
		return e
	}
	return nil
}
