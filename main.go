package main

import (
	"flag"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"log"
	"strings"
	"unicode/utf8"
)

var xlsxPath string
var sheetName string

func init() {
	flag.StringVar(&xlsxPath, "xlsxPath", "", "xlsxPath")
	flag.StringVar(&sheetName, "sheetName", "", "sheetName")
	flag.Parse()
}

func main() {
	if xlsxPath == "" || sheetName == "" {
		log.Fatalln("require xlsxPath or sheetName")
	}
	f, err := excelize.OpenFile(xlsxPath)
	if err != nil {
		log.Fatalln(err)
	}
	rows, err := f.GetRows(sheetName)
	var md string
	for k, v := range rows {
		if len(v) > 0 {
			line := strings.Join(v, " | ")
			md += "| " + line + " |\n"
			if k == 0 {
				md += "| "
				for _, vv := range v {
					count := utf8.RuneCountInString(vv)
					md += strings.Repeat("-", count*2) + " | "
				}
				md += "\n"
			}
		}
	}
	log.Println(md)
}
