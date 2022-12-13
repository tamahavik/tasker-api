package services

type AuthenticationService interface {
	Login(auth string, password string) error
}
