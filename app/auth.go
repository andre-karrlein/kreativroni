package main

import "golang.org/x/crypto/bcrypt"

const hash = "$2a$12$/IGlkEMJ9li8hFTHvLww/.4PR0l.TjKPg5PdguI/aonaOycCrzE1K"

// check the send password
func CheckAuth(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
