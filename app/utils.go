package app

import (
	"errors"
	)

func CleanArgs(name string, flags Flags) (string, error) {
	if len(name) == 0 {
		return "", errors.New("jump: No checkpoint provided")
	}
	
	return name, nil
}