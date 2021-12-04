package requests

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}
