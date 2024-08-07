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

type Options func(*Awk)

func NewAwk(data string) *Awk {

	return &Awk{Data: data}

}

func (a *Awk) DataSplit(splitField, printingField string, columnNumbers ...string) *Awk {
	var fieldsChosen string

	if len(columnNumbers) == 0 {

		splittedData := strings.Split(a.Data, splitField)

		for _, all := range splittedData {
			fieldsChosen += fmt.Sprintf("%v%v", all, printingField)
		}

	} else {
		for _, value := range columnNumbers {

			c := value

			column, err := strconv.Atoi(c)

			if err != nil {

				log.Printf("the string Converter did not work : %v", err)
			}

			splittedData := strings.Split(a.Data, splitField)

			fieldsChosen += fmt.Sprintf("%v%v", splittedData[column], printingField)

		}
	}

	return &Awk{Data: fieldsChosen}
}
