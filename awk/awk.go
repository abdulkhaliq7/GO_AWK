package awk

import (
	"bufio"
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

	reader := strings.NewReader(a.Data)
	scanner := bufio.NewScanner(reader)
	var fieldsChosen strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		if len(chosenColumns) == 0 {
			splittedData := strings.Split(line, splitField)
			for _, all := range splittedData {
				fieldsChosen.WriteString(all)
				fieldsChosen.WriteString(printingField)
			}
			newLine := "\n"
			fieldsChosen.WriteString(newLine)
		} else {
			for _, value := range chosenColumns {
				c := value
				column, err := strconv.Atoi(c)
				if err != nil {
					log.Printf("the string Converter did not work : %v", err)
				}
				splittedData := strings.Split(line, splitField)
				if len(splittedData) > column {
					fieldsChosen.WriteString(splittedData[column])
					fieldsChosen.WriteString(printingField)
				} else {
					log.Printf("the chosen column %v is greater than the length of the string %v", column, len(splittedData))
				}
			}
			newLine := "\n"
			fieldsChosen.WriteString(newLine)
		}
	}
	return &Awk{Data: fieldsChosen.String()}
}

func (a *Awk) Replace(old, new string) *Awk {

	updatedData := strings.ReplaceAll(a.Data, old, new)
	return &Awk{Data: updatedData}
}
