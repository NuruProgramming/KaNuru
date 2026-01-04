package object

import "fmt"

type Error struct {
	Message string
}

func (e *Error) Inspect() string {
	msg := fmt.Sprintf("Kosa: %s", e.Message)
	return msg
}
func (e *Error) Type() ObjectType { return ERROR_OBJ }
