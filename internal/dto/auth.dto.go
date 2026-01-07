package dto

type AuthData struct {
	Email    string
	Password string
}

type AuthResponse struct {
	Status  string
	Message string
	Data    interface{}
	Error   string
}
