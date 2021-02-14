package responses

type TokenGet struct {
	Token   string `json:"token"`
	Profile string `json:"profile"`
	Data    string `json:"data"`
}
