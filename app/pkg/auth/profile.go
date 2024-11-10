package auth

type Profile struct {
	Picture  string `json:"picture"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
}
