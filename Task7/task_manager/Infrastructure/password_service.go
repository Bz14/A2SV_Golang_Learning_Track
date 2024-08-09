package infrastructure

import "golang.org/x/crypto/bcrypt"


type Password struct{
	
}
type PasswordHash interface{
	HashPassword(password string) []byte
	UnHashPassword(existingPassword []byte, newPassword []byte)error
}

func NewPasswordHash()*Password{
	return &Password{}
}

func (passWord *Password)HashPassword(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}
	return hashedPassword
}

func (passWord *Password)UnHashPassword(existingPassword []byte, newPassword []byte) error {
	return bcrypt.CompareHashAndPassword(existingPassword, newPassword)
}