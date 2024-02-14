package tuuid

import (
	"github.com/google/uuid"
)

func FromString(data string) uuid.UUID {
	result, err := uuid.Parse(data)
	if err != nil {
		return uuid.UUID{}
	}
	return result
}
