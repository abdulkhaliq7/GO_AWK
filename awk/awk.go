package awk

import (
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

	if len(chosenColumns) == 0 {

		splittedData := strings.Split(a.Data, splitField)

		for _, all := range splittedData {
			fieldsChosen += fmt.Sprintf("%v%v", all, printingField)
		}

	} else {
		for _, value := range chosenColumns {

			c := value

			column, err := strconv.Atoi(c)

			if err != nil {

				log.Printf("the string Converter did not work : %v", err)
			}

			splittedData := strings.Split(a.Data, splitField)

			// Print the resulting fields
			for i, field := range splittedData {
				if column == i {
					fieldsChosen += fmt.Sprintf("%v%v", field, printingField)
				} else {
					continue
				}
			}

		}
	}

	return &Awk{Data: fieldsChosen}
}

func (a *Awk) Replace(old, new string) *Awk {

	updatedData := strings.ReplaceAll(a.Data, old, new)

	return &Awk{Data: updatedData}
}
