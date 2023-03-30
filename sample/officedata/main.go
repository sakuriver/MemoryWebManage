package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/shakinm/xlsReader/xls"
)

func main() {

	workbook, err := xls.OpenFile("small_1_sheet.xls")

	if err != nil {
		log.Panic(err.Error())
	}

	// Кол-во листов в книге
	// Number of sheets in the workbook
	//
	// for i := 0; i <= workbook.GetNumberSheets()-1; i++ {}

	fmt.Println(workbook.GetNumberSheets())

	sheet, err := workbook.GetSheet(0)

	if err != nil {
		log.Panic(err.Error())
	}

	// Имя листа
	// Print sheet name
	println(sheet.GetName())

	// Вывести кол-во строк в листе
	// Print the number of rows in the sheet
	println(sheet.GetNumberRows())
	records := [][]string{}

	for i := 0; i <= sheet.GetNumberRows(); i++ {
		dataRow := []string{}
		if row, err := sheet.GetRow(i); err == nil {
			for j := 0; j < 3; j++ {
				if cell, err := row.GetCol(j); err == nil {

					// Значение ячейки, тип строка
					// Cell value, string type
					fmt.Println(cell.GetString())

					//fmt.Println(cell.GetInt64())
					//fmt.Println(cell.GetFloat64())

					// Тип ячейки (записи)
					// Cell type (records)
					fmt.Println(cell.GetType())

					// Получение отформатированной строки, например для ячеек с датой или проценты
					// Receiving a formatted string, for example, for cells with a date or a percentage
					xfIndex := cell.GetXFIndex()
					formatIndex := workbook.GetXFbyIndex(xfIndex)
					format := workbook.GetFormatByIndex(formatIndex.GetFormatIndex())
					fmt.Println(format.GetFormatString(cell))
					dataRow = append(dataRow, format.GetFormatString(cell))

				}

			}
			records = append(records, dataRow)
		}

	}

	writeDataFile(records)

}

func writeDataFile(records [][]string) {
	// データ編集チームが設定したExcelからアプリ向けに使える

	f, err := os.Create("data_result.csv") // 書き込む先のファイル
	if err != nil {
		fmt.Println(err)
	}

	w := csv.NewWriter(f)
	w.WriteAll(records) // 一度にすべて書き込む

}
