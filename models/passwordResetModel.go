package models

type PasswordResetRequest struct {
	Username    string `json:"username"`
	NewPassword string `json:"newPassword"`
}
