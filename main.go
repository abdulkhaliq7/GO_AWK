package main

import (
	"fmt"
	"github/abdulkhaliq/awk/awk"
)

func main() {

	data := "633022241|4000000|$30|20240731"

	field := "|"

	printingField := "="

	columnNumber := []string{"1", "3"}

	//printingFiled := ","

	//awk.NewAwk(data, field, printingFiled, columnNumber)

	awk := awk.NewAwk(
		awk.WithDataAndSplitFieldAndPrintingField(data, field, printingField),
		awk.WithColumnNumber(columnNumber),
	)

	output := awk.GetFilteredData()

	fmt.Println(output)
}
