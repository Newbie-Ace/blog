package response

type Login struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expiresAt"`
}
