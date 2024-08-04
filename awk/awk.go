package awk

import (
	_ "fmt"
	_ "log"
	_ "strconv"
	_ "strings"
)

type Awk struct {
	data          string
	splitField    string
	printingField string
	columnNumber  []string
	replaceData   string
}

type Options func(*Awk)

func NewAwk(options ...Options) *Awk {

	a := &Awk{}

	for _, o := range options {
		o(a)
	}

	return a

}

func WithDataAndSplitField(data, field string) Options {
	return func(a *Awk) {
		a.splitField = field
		a.data=data
	}
}

func WithColumnNumber(columnNumber []string) Options {
	return func(a *Awk) {
		a.columnNumber = columnNumber
	}
}


func WithReplaceData(replaceData string) Options {
	return func(a *Awk) {
		a.replaceData=replaceData
	}
}

func WithPrintingField(printingField string) Options {
	return func(a *Awk) {
		a.printingField=printingField
	}
}



/*

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

*/
