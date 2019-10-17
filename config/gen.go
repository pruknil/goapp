package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	excelFileName := "app.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Print(err)
	}
	sheet := xlFile.Sheet
	apiSheet := sheet["API"]
	for irow, row := range apiSheet.Rows {
		if irow > 0 {
			c := row.Cells
			fmt.Println(c[0])
			fmt.Println(c[1])
			fmt.Println(c[2])
			fmt.Println(c[3])
			//for _, cell := range row.Cells {
			//	text := cell.String()
			//	fmt.Printf("%s\n", text)
			//}
		}
	}

	/*	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}*/
}
