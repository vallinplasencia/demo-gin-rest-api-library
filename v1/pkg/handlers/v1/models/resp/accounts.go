package resp

// Login ...
type Login struct {
	AuthTwoFactor bool   `json:"auth_two_factor"`
	Token         *Token `json:"token,omitempty"`
	Fullname      string `json:"fullname,omitempty"`
	Avatar        string `json:"avatar,omitempty"`
	DeviceID      string `json:"device_id,omitempty"`
}

// Token ...
type Token struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
