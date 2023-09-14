package utils

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(password string) (hashPassword string, err error) {
  passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  hashPassword = string(passwordBytes)
  return
}

func CheckPassword(password, comparePasswrod string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(comparePasswrod), []byte(password))
  if err != nil {
    return false
  }
  return true
}
