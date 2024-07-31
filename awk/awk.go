package awk

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Awk(data, field, columnNumber string) {

	//fmt.Println(data, field, columnNumber)

	column, err := strconv.Atoi(columnNumber)

	if err != nil {

		log.Printf("string is not converted to int : %v", err)
	}

	column = column - 1 

	columns := strings.Split(data, "|")

	fmt.Println(columns[column])
}
