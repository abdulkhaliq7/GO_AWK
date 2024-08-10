package main

import (
	"fmt"
	"github/abdulkhaliq/awk/awk"
)

func main() {

<<<<<<< HEAD
	data := `633022241|4000000|$30|20240731
			 633022241|4000000|$30|20240731`
=======
	data := ``
	
	//data = "633022241|4000000|$30|20240731"
>>>>>>> d719d520efa36760f6b1d7addbc7ef41f07aab8f

	newAwk := awk.NewAwk(data)

	output := newAwk.DataSplit("|", " ", "0", "2").Replace("633022241", "634001157")

	fmt.Println(output.Data)

}
