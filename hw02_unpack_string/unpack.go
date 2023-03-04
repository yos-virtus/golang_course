package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(packedString string) (string, error) {
	var prevSymbol rune
	var stringAccumulator strings.Builder

	for _, currentSymbol := range packedString {
		repetitionsCountForSymbol, err := strconv.Atoi(string(currentSymbol))
		if err == nil {
			if prevSymbol == 0 || unicode.IsDigit(prevSymbol) {
				return "", ErrInvalidString
			}

			trimPreviouslyAddedSymbol(&stringAccumulator)

			repeatedSymbolSequence := strings.Repeat(string(prevSymbol), repetitionsCountForSymbol)
			stringAccumulator.WriteString(repeatedSymbolSequence)
		} else {
			stringAccumulator.WriteRune(currentSymbol)
		}
		prevSymbol = currentSymbol
	}

	return stringAccumulator.String(), nil
}

func trimPreviouslyAddedSymbol(stringAccumulator *strings.Builder) {
	if str := stringAccumulator.String(); str != "" {
		stringAccumulator.Reset()
		symbolSequence := []rune(str)
		stringAccumulator.WriteString(string(symbolSequence[:len(symbolSequence)-1]))
	}
}
