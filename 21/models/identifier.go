package models

import (
	"fmt"
	"strings"
)

type Identifier struct {
	id string
}

func CreateIdentifier(id string) Operand {
	return Identifier{strings.TrimSpace(id)}
}

func (id Identifier) GetId() string {
	return id.id
}

func (id Identifier) String() string {
	return fmt.Sprintf("ID[%s]", id.id)
}

func (id Identifier) Resolve(id_resolver func(string) (int64, error)) (int64, error) {
	return id_resolver(id.id)
}

func (thisId Identifier) Equal(other Operand) bool {
	otherId, ok := other.(Identifier)
	return ok && thisId.id == otherId.id
}
