package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-pdf/fpdf"
	"github.com/jlaffaye/ftp"
	"log"
	"net/http"
	"os"
	"time"
)

type Alg2 struct {
	Answer1  string `json:"a"`
	Answer2  string `json:"b"`
	Answer3  string `json:"c"`
	Answer4  string `json:"d"`
	Answer5  string `json:"e"`
	Answer6  string `json:"f"`
	Answer7  string `json:"g"`
	Answer8  string `json:"h"`
	Answer9  string `json:"i"`
	Answer10 string `json:"j"`
	Answer11 string `json:"k"`
	Answer12 string `json:"l"`
	Answer13 string `json:"m"`
	Answer14 string `json:"n"`
	Answer15 string `json:"o"`
	Answer16 string `json:"p"`
	Answer17 string `json:"q"`
	Answer18 string `json:"r"`
	Answer19 string `json:"s"`
	Answer20 string `json:"t"`
	Answer21 string `json:"u"`
	Answer22 string `json:"v"`
	Answer23 string `json:"w"`
	Answer24 string `json:"x"`
	Answer25 string `json:"y"`
	Answer26 string `json:"z"`
	Answer27 string `json:"aa"`
	Answer28 string `json:"ab"`
	Answer29 string `json:"ac"`
	Answer30 string `json:"ad"`
	Answer31 string `json:"ae"`
	Answer32 string `json:"af"`
	Answer33 string `json:"ag"`
	Answer34 string `json:"ah"`
	Answer35 string `json:"ai"`
	Answer36 string `json:"aj"`
	Answer37 string `json:"ak"`
	Answer38 string `json:"al"`
}

type Chem struct {
	Answer1 string `json:"a"`
	Answer2 string `json:"b"`
	Answer3 string `json:"c"`
	Answer4 string `json:"d"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/alg2", func(w http.ResponseWriter, r *http.Request) {
		score := 0
		question := 0
		answers := [38]string{
			"2x^{8}y^{12}\\sqrt{5y}",                            //1
			"2w^4\\sqrt[3]{3}",                                  //2
			"\\textnormal{Identity property of multiplication}", //3
			"-\\frac {w^4} {6}",                                 //4
			"6x(3x^3+2x^2-1)",                                   //5
			"-i",                                                //6
			"-7",                                                //7
			"\\dfrac{-2(3x+2)}{3(x+2)}",                         //8
			"\\dfrac{13x-47}{(x+1)(x-4)}",                       //9
			"6x^{\\frac{9}{2}}y^{\\frac{25}{2}}",                //10
			"x=\\frac{27}{14}",                                  //11
			"1",                                                 //12
			"\\includegraphics[height=7em, width=7em]{/graphs/question13/graph-1.svg}", //13
			"x=-1",                                  //14
			"\\set{(-4,3),(-1,0)}",                  //15
			"|2x+7|<15",                             //16
			"x=(10)",                                //17
			"\\set{-\\frac{3}{2}, \\frac{1}{2}}",    //18
			"\\set{\\frac{8}{3}}",                   //19
			"\\set{(-2,-6),(5,22)}",                 //20
			"A,E",                                   //21
			"-\\infin<x<1",                          //22
			"8x^4-19",                               //23
			"95\\%",                                 //24
			"(1,-2)",                                //25
			"f(x)=3(\\frac{1}{5})^x",                //26
			"8",                                     //27
			"V=kr^2h",                               //28
			"y=-10",                                 //29
			"f(x)\\textnormal{ approaches }\\infin", //30
			"y=23,100(0.85)^x",                      //31
			"\\frac{27}{5}",                         //32
			"\\frac{12}{7}",                         //33
			"g^{-1}(x)=\\sqrt[3]{x-11}",             //34
			"\\set{x|x>1}",                          //35
			"14.1\\textnormal{ grams}",              //36
			"70",                                    //37
			"(20)",                                  //38
		}

		enableCors(&w)

		var alg2 Alg2
		marginCell := 2.

		curtime := time.Now()
		pdfname := curtime.Format("2006102150405") + ".pdf"
		pdf := fpdf.New("P", "mm", "Letter", "")
		pdf.AddPage()
		pdf.SetFont("Arial", "", 12)
		pdf.SetTitle("Algebra 2 Practice SOL Results", true)
		pdf.SetAuthor("SOL Practice Application", true)
		pdf.SetCreationDate(time.Date(curtime.Year(), curtime.Month(), curtime.Day(), curtime.Hour(), curtime.Minute(), curtime.Second(), curtime.Nanosecond(), time.UTC))

		errr := json.NewDecoder(r.Body).Decode(&alg2)
		if errr != nil {
			panic(errr)
			return
		}
		pagew, pageh := pdf.GetPageSize()
		mleft, mright, _, mbottom := pdf.GetMargins()

		cols := []float64{60, 100, pagew - mleft - mright - 100 - 60}
		rows := [][]string{}
		rows = append(rows, []string{"Your Answer", "Correct Answer", ""})

		answered := [38]string{alg2.Answer1, alg2.Answer2, alg2.Answer3, alg2.Answer4, alg2.Answer5, alg2.Answer6, alg2.Answer7, alg2.Answer8, alg2.Answer9, alg2.Answer10, alg2.Answer11, alg2.Answer12, alg2.Answer13, alg2.Answer14, alg2.Answer15, alg2.Answer16, alg2.Answer17, alg2.Answer18, alg2.Answer19, alg2.Answer20, alg2.Answer21, alg2.Answer22, alg2.Answer23, alg2.Answer24, alg2.Answer25, alg2.Answer26, alg2.Answer27, alg2.Answer28, alg2.Answer29, alg2.Answer30, alg2.Answer31, alg2.Answer32, alg2.Answer33, alg2.Answer34, alg2.Answer35, alg2.Answer36, alg2.Answer37, alg2.Answer38}
		for i := 1; i < 38; i++ {
			if answered[i] == answers[i] {
				rows = append(rows, []string{answered[i], answers[i], "Correct"})
				score++
				question++
			} else {
				rows = append(rows, []string{answered[i], answers[i], "Wrong"})
				question++
			}
		}

		for _, row := range rows {
			_, lineHt := pdf.GetFontSize()
			height := lineHt + marginCell

			x, y := pdf.GetXY()
			// add a new page if the height of the row doesn't fit on the page
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

		err := json.NewEncoder(w).Encode(score)
		if err != nil {
			panic(err)
			return
		}
	})

	http.HandleFunc("/chem", func(w http.ResponseWriter, r *http.Request) {
		score := 0
		question := 0
		answers := [4]string{
			"1\\%",
			"\\textnormal{a temperature probe}",
			"\\textnormal{Double-replacement}",
			"(2-)",
		}

		var chem Chem
		marginCell := 2.

		errr := json.NewDecoder(r.Body).Decode(&chem)
		if errr != nil {
			panic(errr)
			return
		}
		enableCors(&w)
		curtime := time.Now()
		pdfname := curtime.Format("2006102150405") + ".pdf"
		pdf := fpdf.New("P", "mm", "Letter", "")
		pdf.AddPage()
		pdf.SetFont("Arial", "", 12)
		pdf.SetTitle("Chemistry Practice SOL Results", true)
		pdf.SetAuthor("SOL Practice Application", true)
		pdf.SetCreationDate(time.Date(curtime.Year(), curtime.Month(), curtime.Day(), curtime.Hour(), curtime.Minute(), curtime.Second(), curtime.Nanosecond(), time.UTC))
		pagew, pageh := pdf.GetPageSize()
		mleft, mright, _, mbottom := pdf.GetMargins()
		cols := []float64{60, 100, pagew - mleft - mright - 100 - 60}
		rows := [][]string{}
		rows = append(rows, []string{"Your Answer", "Correct Answer", ""})

		answered := [4]string{chem.Answer1, chem.Answer2, chem.Answer3, chem.Answer4}
		for i := 1; i < 4; i++ {
			if answered[i] == answers[i] {
				rows = append(rows, []string{answered[i], answers[i], "Correct"})
				score++
				question++
			} else {
				rows = append(rows, []string{answered[i], answers[i], "Wrong"})
				question++
			}
		}

		c, ftperr := ftp.Dial(os.Getenv("FTP_IP"), ftp.DialWithTimeout(5*time.Second))
		if ftperr != nil {
			log.Fatal("ftp failed connection: " + ftperr.Error())
		}

		ftperr = c.Login(os.Getenv("FTP_USER"), os.Getenv("FTP_PASS"))
		if ftperr != nil {
			log.Fatal("ftp failed login: " + ftperr.Error())
		}

		for _, row := range rows {
			curx, y := pdf.GetXY()
			x := curx

			height := 0.
			_, lineHt := pdf.GetFontSize()

			for i, txt := range row {
				lines := pdf.SplitLines([]byte(txt), cols[i])
				h := float64(len(lines))*lineHt + marginCell*float64(len(lines))
				if h > height {
					height = h
				}
			}

			if pdf.GetY()+height > pageh-mbottom {
				pdf.AddPage()
				y = pdf.GetY()
			}
			for i, txt := range row {
				width := cols[i]
				pdf.Rect(x, y, width, height, "")
				pdf.MultiCell(width, lineHt+marginCell, txt, "", "", false)
				x += width
				pdf.SetXY(x, y)
			}
			pdf.SetXY(curx, y+height)
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

		err := json.NewEncoder(w).Encode(score)
		if err != nil {
			panic(err)
			return
		}
	})
	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
