package practice

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyMessage  = errors.New("пустое сообщение")
	ErrNotifierPanic = errors.New("паника канала")
)

type SendError struct {
	Channel string
	Reason  string
}

func (e *SendError) Error() string {
	return fmt.Sprintf("канал %s: %s", e.Channel, e.Reason)
}
