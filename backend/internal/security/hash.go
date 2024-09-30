package security

import (
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

func HashBcrypt(plain string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.MinCost)
	return string(bytes), err
}

func CheckBcryptHash(plain, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	return err == nil
}

func HashSha3256(plain string) string {
	h := sha3.New256()
	h.Write([]byte(plain))
	return hex.EncodeToString(h.Sum(nil))
}

func CheckSha3256Hash(plain, hash string) bool {
	return hash == HashSha3256(plain)
}
