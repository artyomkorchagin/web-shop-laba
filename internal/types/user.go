package types

type User struct {
	UserID       int    `json:"user_id"`
	Username     string `json:"username"`
	DateOfBirth  string `json:"date_of_birth"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	Role         string `json:"role"`
}

type CreateUserRequest struct {
	Username    string `form:"username"`
	Email       string `form:"email"`
	Password    string `form:"password"`
	DateOfBirth string `form:"date_of_birth"`
}

func NewUser(req CreateUserRequest) *User {
	return &User{
		DateOfBirth:  req.DateOfBirth,
		Email:        req.Email,
		PasswordHash: req.Password,
		Username:     req.Username,
		Role:         "customer",
	}
}
