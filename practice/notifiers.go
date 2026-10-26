package practice

type BaseNotifier struct {
	Channel string
}

func (b BaseNotifier) Label(message string) string {
	return "[" + b.Channel + "] " + message
}

// Имя канала доступно через встраивание BaseNotifier.
func (b BaseNotifier) Name() string {
	return b.Channel
}

type EmailNotifier struct {
	BaseNotifier
	Address  string
	Messages []string
}

func NewEmailNotifier(address string) *EmailNotifier {
	return &EmailNotifier{
		BaseNotifier: BaseNotifier{
			Channel: "email",
		},
		Address: address,
	}
}

func (n *EmailNotifier) Send(message string) error {
	if message == "" {
		return ErrEmptyMessage
	}
	if n.Address == "" {
		return &SendError{
			Channel: n.Channel,
			Reason:  "не задан адрес",
		}
	}

	n.Messages = append(n.Messages, n.Label(message))
	return nil
}

type PushNotifier struct {
	BaseNotifier
	DeviceID string
	Messages []string
}

func NewPushNotifier(deviceID string) *PushNotifier{
	return &PushNotifier{
		BaseNotifier: BaseNotifier{
			Channel: "push",
		},
		DeviceID: deviceID,
	}
}

func (n *PushNotifier) Send(message string) error {
	if message == "" {
		return ErrEmptyMessage
	}
	if n.DeviceID == "" {
		return &SendError{
			Channel: n.Channel,
			Reason:  "не задан device id",
		}
	}

	n.Messages = append(n.Messages, n.Label(message))
	return nil
}

type BrokenNotifier struct {
	BaseNotifier
}

func NewBrokenNotifier() *BrokenNotifier {
	return &BrokenNotifier{
		BaseNotifier: BaseNotifier{
			Channel: "broken",
		},
	}
}

func (n *BrokenNotifier) Send(message string) error {
	panic("драйвер канала повреждён")
}
