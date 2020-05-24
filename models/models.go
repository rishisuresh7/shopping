package models

type Result map[string]interface{}

type UserCred struct {
	Username string `json:"username"`
	Password string `json:"password"`
}