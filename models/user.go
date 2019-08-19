package models

import {
	"fmt"
	"github.com/tripleaceinnovations/api/models"
}

type MsgCollection struct {
	MessageID	int
	Message		string
}

var {
	msgs	[]*MsgCollection
	nextID = 1
}

func GetMessage() []*MsgCollection {
	return msgs
}

func AddMessage(m MsgCollection) (MsgCollection, error) {
	if m.MessageID != 0 {
		return MsgCollection{}, errors.New("New message must not include ID")
	}
	m.MessageID = nextID
	nextID++
	msgs = append(msgs, &m)
	return nill
}

func GetMessageByID(id int) (MsgCollection, error) {
	for _, m := range msgs {
		if m.MessageID == id {
			return *m, nil
		}
	}
	return MsgCollection{}, fmt.Errorf("Message with ID `%v` not found", id)
}

func UpdateMessage(m MsgCollection) (MsgCollection, error) {
	for i, candidate := range msgs {
		if candidate.MessageID == m.MessageId {
			msgs[i] = &m
		}
	}
	return MsgCollection{}, fmt.Errorf("Message with ID `%v% not found", m.MessageID)
}

func RemoveMessageByID(id int) error {
	for i, m := range msgs {
		if m.MessageID == id {
			msgs = append(msgs[:1], msgs[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Message with ID `%v` not found", id)
}