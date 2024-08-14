package utils

import (
	"github.com/google/uuid"

	"strings"
)

func UUID() string {
	u := uuid.New()
	return strings.ReplaceAll(u.String(), "-", "")
}
