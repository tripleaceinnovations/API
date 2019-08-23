package models

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"unicode/utf8"
)

type Phrase struct {
	ID                  int
	Message             string
	IsMessagePalindrome bool
}

var (
	phrases       []*Phrase
	nextID        = 1
	verifiedInput string
)

func GetMessages() []*Phrase {
	return phrases
}

func AddMessage(p Phrase) (Phrase, error) {
	if p.ID != 0 {
		return Phrase{}, errors.New("New message must not include ID in the request")
	}
	p.ID = nextID
	if len(strings.TrimSpace(strings.ToUpper(p.Message))) != 0 {
		p.IsMessagePalindrome = IsPalindrome(p.Message)
		nextID++
		phrases = append(phrases, &p)
		return p, nil
	}
	return Phrase{}, errors.New("New message cannot be empty")
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
			if len(strings.TrimSpace(strings.ToUpper(p.Message))) != 0 {
				p.IsMessagePalindrome = IsPalindrome(p.Message)
				phrases[i] = &p
				return p, nil
			}
			return Phrase{}, errors.New("Message to be updated cannot be empty")
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

func RemoveAllMessages() error {
	for i := range phrases {
		phrases = append(phrases[:i], phrases[:i]...)
		log.Println("... deleting all messages. phrases is now:", phrases)
		return nil
	}
	return fmt.Errorf("Message '%v' was not deleted", phrases)
}

func IsPalindrome(input string) bool {
	input = strings.TrimSpace(strings.ToUpper(input))
	verifiedInput = Verify(input)
	if verifiedInput != input {
		log.Println("... input is a not palindrome: reversed input is ", verifiedInput)
		return false
	}
	log.Println(" ... input is a palindrome: reversed input is ", verifiedInput)
	return true
}

func Verify(s string) string {
	inputLength := len(s)
	buffer := make([]byte, inputLength)
	for k := 0; k < inputLength; {
		j, size := utf8.DecodeRuneInString(s[k:])
		k += size
		utf8.EncodeRune(buffer[inputLength-k:], j)
	}
	log.Println("... reversing the message: result is: ", string(buffer))
	return string(buffer)
}
