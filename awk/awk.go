package awk

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
)

type Awk struct {
	Data string
}

func NewAwk(data string) *Awk {

	return &Awk{Data: data}

}

func (a *Awk) DataSplit(splitField, printingField string, chosenColumns ...string) *Awk {

	var wg sync.WaitGroup

	channel := make(chan string, 200)

	/*
		// Create a new reader for the data string
		reader := strings.NewReader(a.Data)

		// Create a scanner to read data line by line
		scanner := bufio.NewScanner(reader)

		// Scan the data line by line

		//fmt.Println(a.Data)

		for scanner.Scan() {
			wg.Add(1)

			line := scanner.Text()

			go filter(line, splitField, printingField, channel, &wg, chosenColumns...)

		}

	*/

	for  {
		if a.Data != "" {
			wg.Add(1)
			go filter(a.Data, splitField, printingField, channel, &wg, chosenColumns...)
		}else {
			break 
		}
	}

	go func() {
		wg.Wait()
		close(channel)
	}()

	for output := range channel {
		return &Awk{Data: output}
	}

	return nil
}

func filter(line, splitField, printingField string, channel chan<- string, wg *sync.WaitGroup, chosenColumns ...string) {

	var fieldsChosen string

	defer wg.Done()

	if len(chosenColumns) == 0 {

		splittedData := strings.Split(line, splitField)

		for _, all := range splittedData {
			fieldsChosen += fmt.Sprintf("%v%v", all, printingField)
		}

		newLine := "\n"

		fieldsChosen += fmt.Sprintf("%v", newLine)

	} else {
		for _, value := range chosenColumns {

			c := value

			column, err := strconv.Atoi(c)

			if err != nil {

				log.Printf("the string Converter did not work : %v", err)
			}

			splittedData := strings.Split(line, splitField)

			if len(splittedData) > column {
				fieldsChosen += fmt.Sprintf("%v%v", splittedData[column], printingField)
			} else {

				log.Printf("the chosen column %v is greater than the length of the string %v", column, len(splittedData))
			}

		}

		newLine := "\n"

		fieldsChosen += fmt.Sprintf("%v", newLine)
	}

	channel <- fieldsChosen
}

func (a *Awk) Replace(old, new string) *Awk {

	updatedData := strings.ReplaceAll(a.Data, old, new)

	return &Awk{Data: updatedData}
}
