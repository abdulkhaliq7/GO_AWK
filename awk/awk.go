package awk

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Awk(data, splitField, printingField string, columnNumber []string) {

	var fieldsChosen string

	for _, value := range columnNumber {

		c := value

		column, err := strconv.Atoi(c)

		if err != nil {

			log.Printf("the string Converter did not work : %v", err)
		}

		columnChose := strings.Split(data, splitField)

		//this is to print if the entered column is 0
		if column == 0 {

			for _, all := range columnChose {

				fieldsChosen += fmt.Sprintf("%v%v", all, printingField)

			}

			continue
		}

		column = column - 1

		fieldsChosen += fmt.Sprintf("%v%v", printingField, columnChose[column])

	}

	fmt.Println(fieldsChosen)
}
