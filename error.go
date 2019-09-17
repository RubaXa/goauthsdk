package oauth

import (
	"encoding/json"
	"fmt"
)

type ErrorResponse struct {
	Type        string `json:"error"`
	Code        int    `json:"error_code"`
	Description string `json:"error_description"`
}

func (e *ErrorResponse) String() string {
	return fmt.Sprintf("[%s] %s", e.Type, e.Description)
}

func (e *ErrorResponse) Error() string {
	return e.String()
}

func ParseErrorResponse(b []byte) error {
	err := &ErrorResponse{}
	_ = json.Unmarshal(b, err)
	if err.Code > 0 {
		return err
	}

	return nil
}