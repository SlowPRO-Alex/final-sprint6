package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/SlowPRO-Alex/final-sprint6/tree/first-iteration/pkg/morse"
)

var output string

func TypeDetector(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("Ошибка: строка пуста!")
	}
	fmt.Println("Определяем тип данных...")
	textType := "текст"
	for _, v := range input {
		vstr := fmt.Sprintf("%c", v)

		if strings.ContainsAny(vstr, "-.") {
			textType = "код морзе"
			break
		}
	}
	if textType == "код морзе" {
		output = morse.ToText(input)
	} else {
		output = morse.ToMorse(input)
	}
	fmt.Printf("%s это %s\n", input, textType)
	return output, nil
}
