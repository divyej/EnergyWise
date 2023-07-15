package utils

import (
	"encoding/json"
	"io"
)

//  ParseJson parses the io.ReadCloser to respected type, usually used with models
//
// Always pass the pointer

func ParseJson(body io.ReadCloser, v any) error {

	if body, err := io.ReadAll(body); err == nil {
		if err := json.Unmarshal(body, v); err == nil {
			return err
		}
	}
	return nil
}
