package model

type ConfigJsonBody struct {
	Id string `json:"id"`
	CaptchaType string `json:"captcha_type"`
	VerifyValue string `json:"verify_value"`
}