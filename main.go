package main

import (
	"fmt"
	"github/abdulkhaliq/awk/awk"
	"log"
	"os"
	"runtime/pprof"
)

func main() {

	// Start CPU profiling
	f, err := os.Create("cpu.pprof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	//files read
	data := `633022241|4000000|$30|20240731
	633022241|4000000|$30|20240731`

	newAwk := awk.NewAwk(data)

	output := newAwk.DataSplit("|", " ", "0", "2").Replace("633022241", "634001157")

	//output := newAwk.DataSplit("|", " ").Replace("633022241", "634001157")

	fmt.Println(output.Data)

}
