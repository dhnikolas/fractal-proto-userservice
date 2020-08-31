package userservice
import (
	"encoding/base64"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	jsoniter "github.com/json-iterator/go"
)

func (x *User) ValidateSave () error {
	return validation.ValidateStruct(x,
		validation.Field(&x.Email, validation.Required, is.Email),
		validation.Field(&x.FirstName, validation.Required),
		validation.Field(&x.LastName, validation.Required),
		validation.Field(&x.Password, validation.Required),
	)
}

func (x *User) GetIdToken () string {
	userBody, _ := jsoniter.Marshal(x)
	userToken := base64.RawURLEncoding.EncodeToString(userBody)

	return userToken
}
