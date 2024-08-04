package main

import (
	"fmt"
	"github/abdulkhaliq/awk/awk"
	"log"
)

func main() {

	data := "633022241|4000000|$30|20240731"

	field := "|"

	printingField := "="

	//columnNumber := []string{"1", "3"}

	//printingFiled := ","

	//awk.NewAwk(data, field, printingFiled, columnNumber)

	awk := awk.NewAwk(
		awk.WithDataAndSplitFieldAndPrintingField(data, field, printingField),
		//awk.WithColumnNumber(columnNumber),
		awk.WithReplaceAll("633022241", "634001157"),
	)

	output, err := awk.GetFilteredData()

	if err != nil {
		log.Println(err)
	}

	fmt.Println(output)
}
