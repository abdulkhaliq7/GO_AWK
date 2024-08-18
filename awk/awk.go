package awk

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Awk struct {
	Data string
}

func NewAwk(data string) *Awk {

	return &Awk{Data: data}

}

func (a *Awk) DataSplit(splitField, printingField string, chosenColumns ...string) (*Awk, error) {

	reader := strings.NewReader(a.Data)
	scanner := bufio.NewScanner(reader)
	var fieldsChosen strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		splittedData := strings.Split(line, splitField)
		if len(chosenColumns) == 0 {
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
					return nil, fmt.Errorf("column conversion error: %v", err)
				}
				if len(splittedData) > column {
					fieldsChosen.WriteString(splittedData[column])
					fieldsChosen.WriteString(printingField)
				} else {
					return nil, fmt.Errorf("chosen column %v is out of bounds (line length: %v)", column, len(splittedData))
				}
			}
			newLine := "\n"
			fieldsChosen.WriteString(newLine)
		}
	}
	return &Awk{Data: fieldsChosen.String()}, nil
}

func (a *Awk) Replace(old, new string) *Awk {

	updatedData := strings.ReplaceAll(a.Data, old, new)
	return &Awk{Data: updatedData}
}
