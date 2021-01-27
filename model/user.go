package model

type (
	User struct {
		Username string `json:"username" form:"username"`
		Password string `json:"password" form:"password"`
		Email    string `json:"email,omitempty"`
		ID       string `json:"id,omitempty"`
	}
)
