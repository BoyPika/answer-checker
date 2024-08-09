package main

import (
	"fmt"
	"github.com/go-pdf/fpdf"
	"github.com/jlaffaye/ftp"
	"log"
	"os"
	"time"
)

func genpdf(title string, rowsa [][]string) {
	marginCell := 2.
	curtime := time.Now()
	pdfname := curtime.Format("2006102150405") + ".pdf"
	pdf := fpdf.New("P", "mm", "Letter", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.SetTitle(title+" Practice SOL Results", true)
	pdf.SetAuthor("SOL Practice Application", true)
	pdf.SetCreationDate(time.Date(curtime.Year(), curtime.Month(), curtime.Day(), curtime.Hour(), curtime.Minute(), curtime.Second(), curtime.Nanosecond(), time.UTC))
	pagew, pageh := pdf.GetPageSize()
	mleft, mright, _, mbottom := pdf.GetMargins()
	cols := []float64{60, 100, pagew - mleft - mright - 100 - 60}
	rows := [][]string{}
	rows = append(rows, []string{"Your Answer", "Correct Answer", ""})

	rown := 0
	for _, r := range rowsa {
		rows = append(rows, []string{r[0], r[1], r[2]})
		rown++
	}

	for _, row := range rows {
		_, lineHt := pdf.GetFontSize()
		height := lineHt + marginCell

		x, y := pdf.GetXY()
		if y+height >= pageh-mbottom {
			pdf.AddPage()
			x, y = pdf.GetXY()
		}
		if row[2] == "Correct" {
			pdf.SetFillColor(95, 247, 64)
		} else if row[2] == "Wrong" {
			pdf.SetFillColor(255, 69, 56)
		} else {
			pdf.SetFillColor(255, 255, 255)
		}
		for i, txt := range row {
			width := cols[i]
			pdf.Rect(x, y, width, height, "FD")
			pdf.ClipRect(x, y, width, height, false)
			pdf.Cell(width, height, txt)
			pdf.ClipEnd()
			x += width
		}
		pdf.Ln(-1)
	}
	c, ftperr := ftp.Dial(os.Getenv("FTP_IP"), ftp.DialWithTimeout(5*time.Second))
	if ftperr != nil {
		log.Fatal("ftp failed connection: " + ftperr.Error())
	}

	ftperr = c.Login(os.Getenv("FTP_USER"), os.Getenv("FTP_PASS"))
	if ftperr != nil {
		log.Fatal("ftp failed login: " + ftperr.Error())
	}

	pdferr := pdf.OutputFileAndClose(pdfname)
	if pdferr != nil {
		panic(pdferr)
	}

	localFile, ftperr := os.Open(pdfname)

	ftperr = c.ChangeDir("/htdocs")
	ftperr = c.Stor(pdfname, localFile)

	filecloseerr := localFile.Close()
	if filecloseerr != nil {
		panic(filecloseerr)
	}

	if ftperr := c.Quit(); ftperr != nil {
		log.Fatal(ftperr)
	}

	oserr := os.Remove(pdfname)
	if oserr != nil {
		panic(oserr)
	}
	fmt.Println("PDF Generated as " + pdfname)
}
