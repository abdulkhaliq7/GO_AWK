package main

import (
	"bufio"
	"fmt"
	"github/abdulkhaliq/awk/awk"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {

	// Get the current time
	start := time.Now()

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

	file, err := os.Open("cdr.txt")

	if err != nil {
		log.Println(err)
	}

	// Create a buffer to read the file
	bufferSize := 10485760
	// Read strings using bufio.Reader
	for {

		reader := bufio.NewReaderSize(file, bufferSize)

		// Read until newline and store in a buffer
		line, err := reader.ReadString('\n')

		fmt.Println(line)
		if err != nil {
			if err.Error() != "EOF" {
				log.Fatal(err)
			}
		}

		newAwk := awk.NewAwk(line)

		output := newAwk.DataSplit("|", "", "35").DataSplit(`\`, "", "26", "70")

		fmt.Println(output.Data)
	}

	fmt.Println("Total time taken to complete The whole process : ", time.Since(start))

}
