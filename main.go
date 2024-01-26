package main

import (
	"content_automater/GoSheet"
)

func main() {
	Sheet, err := GoSheet.OpenFile("keyword.xlsx")
	if err != nil {
		panic(err)
	}
	defer GoSheet.CloseFile(Sheet)
}
