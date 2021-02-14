package requests

type TokenGet struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}
