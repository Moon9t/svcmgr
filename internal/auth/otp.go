package auth

import (
	"bytes"
	"image/png"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type OTPService struct {
	issuer string
}

func NewOTPService(issuer string) *OTPService {
	return &OTPService{issuer: issuer}
}

func (o *OTPService) GenerateSecret(user string) (*otp.Key, error) {
	return totp.Generate(totp.GenerateOpts{
		Issuer:      o.issuer,
		AccountName: user,
	})
}

func (o *OTPService) ValidateCode(secret string, code string) bool {
	return totp.Validate(code, secret)
}

func (o *OTPService) GetQRCode(secret string) ([]byte, error) {
	key, err := otp.NewKeyFromURL(secret)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		return nil, err
	}

	err = png.Encode(&buf, img)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
