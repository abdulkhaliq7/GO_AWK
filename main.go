package main

import (
	"fmt"
	"github/abdulkhaliq/awk/awk"
)

func main() {

	data := `633022241|4000000|$30|20240731
			 633022241|4000000|$30|20240731`


	newAwk := awk.NewAwk(data)

	output := newAwk.DataSplit("|", " ", "0", "2").Replace("633022241", "634001157")

	fmt.Println(output.Data)

}
