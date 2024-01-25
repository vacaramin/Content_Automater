package GoSheet

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func OpenFile(filename string) (file *excelize.File, err error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return f, nil
}

func CloseFile(file *excelize.File) (err error) {
	if err := file.Close(); err != nil {
		return err
	}
	return nil
}
func GetAllSheetNames(filename string) (file *excelize.File, err error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return f, nil
}
