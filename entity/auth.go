package entity

import "golang.org/x/crypto/bcrypt"

type Auth struct {
	ID       string
	Email    string
	Password string
}

func (a Auth) IsExists() bool {
	return a != Auth{}
}

func (a Auth) VerifyPassword(plainPassword string) error {
	if a.Password != plainPassword {
		return ErrInvalidPassword
	}
	return nil
}

func (a *Auth) GeneratePassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	a.Password = string(hash)
	return nil
}

func (a Auth) GenerateSession() (Session, error) {
	return Session{
		ID:    a.ID,
		Token: "User Token",
	}, nil
}

type Session struct {
	ID    string
	Token string
}
