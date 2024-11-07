package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

//Hello devuelve un saludo para la persona especifica

func Hello(name string) (string, error) {
	//Devuelve un saludo que incluye el nomrbe de un mensaje.
	if name == "" {
		return "", errors.New("Nombre vacio")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hola %v! !Bienvenido!",
		"¡Que bueno verte, %v!",
		"Saludo, %v! ¡Encantada de conocerte!",
	}
	return formats[rand.Intn(len(formats))]
}
