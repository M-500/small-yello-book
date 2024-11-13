package utils

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"strings"
)

// GeneratePwd
// @Description:
// @param plaintext
// @return ciphertext
// @return err
func GeneratePwd(plaintext string) (ciphertext string, err error) {
	if len(strings.TrimSpace(plaintext)) == 0 {
		return "", errors.New("plaintext cannot be empty")
	}
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(plaintext, options)
	ciphertext = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	return
}

// CheckPwd
// @Description:
// @return res
func CheckPwd(oldPassword, checkPassword string) bool {
	//options := &password.Options{16, 100, 32, sha512.New}
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	if !strings.Contains(oldPassword, "$") {
		return false
	}
	passwordInfo := strings.Split(oldPassword, "$")
	if len(passwordInfo) < 4 {
		return false // or handle the error more explicitly
	}
	return password.Verify(checkPassword, passwordInfo[2], passwordInfo[3], options)
	//passwordInfo := strings.Split(oldPassword, "$")
	//return password.Verify(checkPassword, passwordInfo[2], passwordInfo[3], options)
}
