package v3

type Messenger struct {
}

type Screen struct {
}

func NewMessenger() *Messenger {
	return &Messenger{}
}

// Photo формирует кадр состояния мира и агентов (Screen) и отдаёт его на отправку пользователю.
func (m *Messenger) Photo() *Screen {
	return &Screen{}
}
