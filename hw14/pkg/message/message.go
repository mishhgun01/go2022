package message

import "strconv"

type Message struct {
	ID   int
	Text string
}

func (m *Message) String() string {
	s := strconv.Itoa(m.ID) + ":" + m.Text
	return s
}
