package main

import (
	"fmt"
	"os"
)

func FindStructsInFile(filename string) ([]string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}
	return FindStructs(bytes)
}

func FindStructs(bytes []byte) ([]string, error) {
	structs := []string{}

	var w, whichType, openBracket, structContent string
	for i := 0; i < len(bytes); {
		start := i
		i, w = readWord(i, bytes)
		if w == "type" {
			i, _ = readWord(i, bytes)
			i, whichType = readWord(i, bytes)
			if whichType == "struct" {
				i, openBracket = readWord(i, bytes)
				// next word should be {
				if openBracket != "{" {
					return structs, fmt.Errorf("invalid struct at position %d", i)
				}
				brackets := 1
				for brackets > 0 {
					i, structContent = readWord(i, bytes)
					if structContent == "{" {
						brackets++
					} else if structContent == "}" {
						brackets--
					}
				}
				structs = append(structs, string(bytes[start:i-1]))
			}
		}
	}

	return structs, nil
}

func readWord(i int, b []byte) (int, string) {
	for i < len(b) && empty(b[i]) {
		i++
	}
	start := i
	for i < len(b) && !empty(b[i]) {
		i++
	}
	return i + 1, string(b[start:i])
}

func empty(b byte) bool {
	switch b {
	case ' ', '	', '\n':
		return true
	default:
		return false
	}
}
