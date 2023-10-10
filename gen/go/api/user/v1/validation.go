package userv1

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (r *RegisterUserRequest) Validate() error {
	err := validation.ValidateStruct(r,
		validation.Field(&r.FirstName, validation.Required, is.Alpha),
		validation.Field(&r.LastName, validation.Required, is.Alpha),
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(5, 100)),
	)
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return nil
}

func (r *LoginUserRequest) Validate() error {
	err := validation.ValidateStruct(r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(5, 100)),
	)
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return nil
}

func (r *LogoutUserRequest) Validate() error {
	err := validation.ValidateStruct(r,
		validation.Field(&r.Uuid, validation.Required, is.UUID),
	)
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return nil
}

func (r *GetUserAccountRequest) Validate() error {
	err := validation.ValidateStruct(r,
		validation.Field(&r.Uuid, validation.Required, is.UUID),
	)
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return nil
}
