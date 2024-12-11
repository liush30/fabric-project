package tool

import (
	"github.com/google/uuid"
	"strings"
)

func GenerateUUIDWithoutDashes() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
