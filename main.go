package main

import "github/abdulkhaliq/awk/awk"

func main() {


	data := "633022241|4000000|$30|20240731"

	field := "|"

	columnNumber := []string{"1", "3", "0"}

	printingFiled := ","

	awk.NewAwk(data, field, printingFiled, columnNumber)
}
