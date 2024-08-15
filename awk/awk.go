package awk

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Awk struct {
	Data string
}

func NewAwk(data string) *Awk {

	return &Awk{Data: data}

}

func (a *Awk) DataSplit(splitField, printingField string, chosenColumns ...string) *Awk {
	var fieldsChosen string

	// Create a new reader for the data string
	reader := strings.NewReader(a.Data)

	// Create a scanner to read data line by line
	scanner := bufio.NewScanner(reader)

	// Scan the data line by line

	for scanner.Scan() {

		line := scanner.Text()

		filter(line, splitField, printingField, &fieldsChosen, chosenColumns...)

	}

	return &Awk{Data: fieldsChosen}
}

func filter(line, splitField, printingField string, fieldsChosen *string, chosenColumns ...string) {
	if len(chosenColumns) == 0 {

		splittedData := strings.Split(line, splitField)

		for _, all := range splittedData {
			*fieldsChosen += fmt.Sprintf("%v%v", all, printingField)
		}

		newLine := "\n"

		*fieldsChosen += fmt.Sprintf("%v", newLine)

	} else {
		for _, value := range chosenColumns {

			c := value

			column, err := strconv.Atoi(c)

			if err != nil {

				log.Printf("the string Converter did not work : %v", err)
			}

			splittedData := strings.Split(line, splitField)

			if len(splittedData) > column {
				*fieldsChosen += fmt.Sprintf("%v%v", splittedData[column], printingField)
			} else {

				log.Printf("the chosen column %v is greater than the length of the string %v", column, len(splittedData))
			}

		}

		newLine := "\n"

		*fieldsChosen += fmt.Sprintf("%v", newLine)
	}
}

func (a *Awk) Replace(old, new string) *Awk {

	updatedData := strings.ReplaceAll(a.Data, old, new)

	return &Awk{Data: updatedData}
}
