package main

import "github/abdulkhaliq/awk/awk"

func main() {


	data := "633022241|4000000|$30|20240731"

	field := "|"

	columnNumber := "3"

	awk.Awk(data, field, columnNumber)
}
