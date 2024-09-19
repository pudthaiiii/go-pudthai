package recaptcha

import (
	"encoding/json"
	"fmt"
	"go-ibooking/internal/infrastructure/logger"
	"io/ioutil"
	"net/http"
	"os"
)

type RecaptchaProvider struct {
	SecretKey string
}

func NewRecaptchaProvider() *RecaptchaProvider {
	return &RecaptchaProvider{
		SecretKey: os.Getenv("GOOGLE_RECAPTCHA_SECRET_KEY"),
	}
}

func (r *RecaptchaProvider) VerifyToken(token string) (bool, error) {
	resp, err := http.PostForm(
		"https://www.google.com/recaptcha/api/siteverify",
		map[string][]string{
			"secret":   {r.SecretKey},
			"response": {token},
		},
	)

	if err != nil {
		logger.Log.Err(err).Msg("failed to verify reCAPTCHA token")
		return false, fmt.Errorf("failed to verify reCAPTCHA token: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Err(err).Msg("failed to read reCAPTCHA response body")
		return false, fmt.Errorf("failed to read reCAPTCHA response body: %w", err)
	}

	var result struct {
		Success bool `json:"success"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		logger.Log.Err(err).Msg("failed to unmarshal reCAPTCHA response")
		return false, fmt.Errorf("failed to unmarshal reCAPTCHA response: %w", err)
	}

	return result.Success, nil
}
