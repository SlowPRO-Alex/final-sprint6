package service

import (
	"fmt"
	"strings"
	"github.com/SlowPRO-Alex/final-sprint6/tree/main/pkg/morse"
)

func typeDetector (input string) (output string) {
	if len(input) = 0 {
		return "Ошибка: строка пуста!"
	}
	fmt.Println("Определяем тип данных...")
	textType := "код морзе"
	for _, v := range input {
		if strings.ContainsAny(v, "-.") != true {
			textType = "текст"
			break
		}
	}
	if textType = "код морзе" {
		output = ToText(input) 
	} else {
		output = ToMorse(input)
	}
	fmt.Println("%s это %s\n", input, textType)
	return output
}