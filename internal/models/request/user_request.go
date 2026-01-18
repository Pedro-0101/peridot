package request

type UserRequest struct {
	Name  string `json:"username"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
}
