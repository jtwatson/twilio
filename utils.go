package twilio

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

func decodeError(r io.Reader) error {
	apiError := &APIError{}
	if err := json.NewDecoder(r).Decode(apiError); err != nil {
		return errors.WithMessage(err, "decodeError()")
	}

	return apiError
}
