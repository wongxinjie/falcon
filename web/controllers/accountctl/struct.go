package accountctl

type MobileLoginRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Via      string `json:"via"`
}

func (mr *MobileLoginRequest) IsValid() bool {
	if len(mr.Mobile) == 0 || len(mr.Password) == 0 || len(mr.Via) == 0 {
		return false
	}
	return true
}

type LoginResponse struct {
	Token string `json:"token"`
}

type AccountResponse struct {
	UserID       int64  `json:"user_id"`
	UserName     string `json:"user_name"`
	RegisteredAt int64  `json:"registered_at"`
	Level        int    `json:"level"`
}
