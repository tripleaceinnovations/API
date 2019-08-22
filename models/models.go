package models

import (
	"errors"
	"fmt"
	"log"
)

type Phrase struct {
	ID                  int
	Message             string
	IsMessagePalindrome bool
}

var (
	phrases       []*Phrase
	nextID        = 1
	reversedinput string
)

func GetMessages() []*Phrase {
	return phrases
}

func AddMessage(p Phrase) (Phrase, error) {
	if p.ID != 0 {
		return Phrase{}, errors.New("New message must not include ID in the request")
	}
	p.ID = nextID
	if len(p.Message) != 0 {
		p.IsMessagePalindrome = IsPalindrome(p.Message)
		nextID++
		phrases = append(phrases, &p)
		return p, nil
	} else {
		return Phrase{}, errors.New("New message cannot be empty")
	}
}

func GetMessageByID(id int) (Phrase, error) {
	for _, p := range phrases {
		if p.ID == id {
			return *p, nil
		}

	}
	log.Println("... GetMessageByID: Validation failed for message ID: ")
	return Phrase{}, fmt.Errorf("Message with ID '%v' not found", id)
}

func UpdateMessage(p Phrase) (Phrase, error) {
	for i, candidate := range phrases {
		if candidate.ID == p.ID {
			if len(p.Message) != 0 {
				p.IsMessagePalindrome = IsPalindrome(p.Message)
				phrases[i] = &p
				return p, nil
			} else {
				return Phrase{}, errors.New("Message to be updated cannot be empty")
			}
		}
	}
	return Phrase{}, fmt.Errorf("Message with ID '%v' not found", p.ID)
}

func RemoveMessageByID(id int) error {
	for i, p := range phrases {
		if p.ID == id {
			phrases = append(phrases[:i], phrases[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Message with ID '%v' not found", id)
}

func IsPalindrome(input string) bool {
	reversedinput = reverse(input)
	if reversedinput != input {
		log.Println("... input is a not palindrome: reversed input is ", reversedinput)
		return false
	} else {
		log.Println(" ... input is a palindrome: reversed input is ", reversedinput)
		return true
	}
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	log.Println("... reversing the message: result is: ", string(chars))
	return string(chars)
}
