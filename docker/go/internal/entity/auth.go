package entity

type Auth struct {
	User
	Accesstoken string `json:"access_token"`
}

type AuthResponse struct {
	UserResponse
	Accesstoken string `json:"access_token"`
}
