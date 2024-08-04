package awk

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Awk struct {
	data string
	splitField string
	printingField string
	columnNumber []string
	replaceData string
}

type Options func(*Awk)


func NewAwk(options ...Options) *Awk {

	a := &Awk{}

	for _, o:= range options {
		o(a)
	}

	return a

}



func NewAwk(options ...Options) {

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
