package entity

type AuthResponse struct {
	UserResponse
	Accesstoken string `json:"access_token"`
}
