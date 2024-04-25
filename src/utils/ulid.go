package utils

import (
	"github.com/oklog/ulid/v2"
)

func GenerateULID() string {
	id := ulid.Make()
	return id.String()
}
