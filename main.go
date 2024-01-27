package main

import (
	"content_automater/GoSheet"
)

func main() {
	Sheet := GoSheet.OpenFile("keyword.xlsx") // Error handled in Gosheet package

	defer GoSheet.CloseFile(Sheet)

}
