package main

import (
	"fmt"
	"github/abdulkhaliq/awk/awk"
)

func main() {

	data := ``

	//data = "633022241|4000000|$30|20240731"

	newAwk := awk.NewAwk(data)

	output := newAwk.DataSplit("|", " ", "35").DataSplit(`\`, "=", "26", "70").DataSplit("=", " ").Replace(",", "")
	fmt.Println(output.Data)

}
