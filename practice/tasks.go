package practice

import (
	"fmt"
)

type Notifier interface {
	Send(message string) error
}

// Дополнительный контракт только для подписи ошибки.
// Основной интерфейс Notifier остаётся прежним.
type namedNotifier interface {
	Name() string
}

func channelName(n Notifier) string {
	if named, ok := n.(namedNotifier); ok {
		return named.Name()
	}
	return fmt.Sprintf("%T", n)
}

func Dispatch(all []Notifier, message string) error {
	for _, n := range all {
		if err := n.Send(message); err != nil {
			return fmt.Errorf("канал %s: %w", channelName(n), err)
		}
	}
	return nil
}

func SafeSend(n Notifier, message string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%w: %v", ErrNotifierPanic, r)
		}
	}()

	return n.Send(message)
}

func SafeDispatch(all []Notifier, message string) []error {
	var errs []error
	for _, n := range all {
		if err := SafeSend(n, message); err != nil {
			errs = append(errs, fmt.Errorf("канал %s: %w", channelName(n), err))
		}
	}
	return errs
}

func EmailAddress(n Notifier) (string, bool) {
	email, ok := n.(*EmailNotifier)
	if !ok || email == nil {
		return "", false
	}
	return email.Address, true
}

func Describe(n Notifier) string {
	switch v := n.(type) {
	case *EmailNotifier:
		if v == nil {
			return "unknown"
		}
		return "email: " + v.Address
	case *PushNotifier:
		if v == nil {
			return "unknown"
		}
		return "push: " + v.DeviceID
	default:
		return "unknown"
	}
}

var _ Notifier = (*EmailNotifier)(nil)

var _ Notifier = (*PushNotifier)(nil)
