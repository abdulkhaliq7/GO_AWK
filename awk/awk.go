package awk

import (
	"fmt"
	"log"
	"strconv"
	"strings"
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

func WithDataAndSplitFieldAndPrintingField(data, field, printingField string) Options {
	return func(a *Awk) {
		a.splitField = field
		a.data = data
		a.printingField = printingField
	}
}

func WithColumnNumber(columnNumber []string) Options {
	return func(a *Awk) {
		a.columnNumber = columnNumber
	}
}

func WithReplaceData(replaceData string) Options {
	return func(a *Awk) {
		a.replaceData = replaceData
	}
}

func WithPrintingField(printingField string) Options {
	return func(a *Awk) {
		a.printingField = printingField
	}
}

func dataAndSplitFieldAndPrintingField(a *Awk) string {
	var fieldsChosen string
	columnChose := strings.Split(a.data, a.splitField)

	for _, all := range columnChose {

		fieldsChosen += fmt.Sprintf("%v%v", all, a.printingField)

	}

	return fieldsChosen
}

func columnNumber(a *Awk) string {
	var fieldsChosen string

	for _, value := range a.columnNumber {

		c := value

		column, err := strconv.Atoi(c)

		if err != nil {

			log.Printf("the string Converter did not work : %v", err)
		}

		columnChose := strings.Split(a.data, a.splitField)

		column = column - 1

		fieldsChosen += fmt.Sprintf("%v%v", a.printingField, columnChose[column])

	}

	return fieldsChosen
}

func (a *Awk) GetFilteredData() string {

	var filteredData string

	if len(a.columnNumber) != 0 {
		filteredData = columnNumber(a)
	} else {
		filteredData = dataAndSplitFieldAndPrintingField(a)
	}

	return filteredData
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
