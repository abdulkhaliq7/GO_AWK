package main

import (
	"bufio"
	"fmt"
	"github/abdulkhaliq/awk/awk"
	"log"
	"os"
	"time"
)

func main() {

	// Get the current time
	start := time.Now()

	file, err := os.Open("cdr.txt")

	if err != nil {
		log.Println(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Print each line
		line := scanner.Text()
		newAwk := awk.NewAwk(line)

		output := newAwk.DataSplit("|", "", "35").DataSplit(`\`, "", "26", "70")

		fmt.Println(output.Data)
	}
	fmt.Println("Total time taken to complete The whole process : ", time.Since(start))

}




