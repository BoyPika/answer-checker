package main

import (
	"encoding/json"
	"net/http"
)

type Grd8mth struct {
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
	Answer39 string `json:"am"`
	Answer40 string `json:"an"`
	Answer41 string `json:"ao"`
	Answer42 string `json:"ap"`
	Answer43 string `json:"aq"`
	Answer44 string `json:"ar"`
	Answer45 string `json:"as"`
	Answer46 string `json:"at"`
	Answer47 string `json:"au"`
	Answer48 string `json:"av"`
	Answer49 string `json:"aw"`
	Answer50 string `json:"ax"`
	Name     string `json:"name"`
}

func grd8mth() {
	http.HandleFunc("/grd8mth", func(w http.ResponseWriter, r *http.Request) {
		score := 0
		question := 0
		answers := [50]string{
			"256",        //1
			"(-18)",      //2
			"12",         //3
			"10\\%",      //4
			"\\sqrt{17}", //5
			"B,D",        //6
			"\\$9.90",    //7
			"3.1*10^{1}, \\dfrac{10}{11}, 52\\%, 0.45", //8
			"\\$42.71", //9
			"\\sqrt{42} \\textnormal{ and } \\sqrt{48}", //10
			"(44)", //11 TODO: Add include graphics
			"-4.8", //12
			"\\textnormal{Between }7\\textnormal{ and }8", //13
			"9",                                   //14
			"5^{2}+x^{2}=12^{2}",                  //15
			"96 \\textnormal{ sq cm}",             //16
			"b",                                   //17 TODO: Add include graphics
			"97\\deg \\textnormal{ and } 83\\deg", //18
			"answer",                              //19 TODO: Fix the frontend for this
			"\\textnormal{The volume is } 10 \\textnormal{ times the volume of the first prism.}", //20
			"44\\textnormal{in}",            //21
			"226\\textnormal{ cu in.}",      //22
			"\\textnormal{Translation}",     //23
			"a",                             //24 TODO: Add include graphics
			"8.7\\textnormal{ m}",           //25
			"408\\textnormal{ sq ft}",       //26
			"A,C",                           //27
			"B,F",                           //28
			"-\\dfrac{3}{2}",                //29
			"c",                             //30 TODO: Add include graphics
			"A,D",                           //31
			"\\$9.36",                       //32
			"9",                             //33
			"a",                             //34 TODO: Add include graphics
			"\\dfrac{1}{4}",                 //35
			"d",                             //36 TODO: Add include graphics
			"a",                             //37 TODO: Fix the frontend for this
			"\\textnormal{No relationship}", //38
			"n=28",                          //39
			"\\textnormal{Between }10\\textnormal{ and } 11\\textnormal{ years}", //40
			"d",              //41 TODO: Add include graphics
			"\\set{1,2,4,5}", //42
			"-3\\geq{x}",     //43
			"13",             //44
			"c",              //45 TODO: Fix the frontend for this
			"\\textnormal{The total number of text messages sent by Kala was the same as the total number sent by Paul for these four days.}", //46
			"\\textnormal{As Marvin's age increased, the time it took him to run the mile decreased.}",                                        //47
			"\\set{(5,1), (3,4), (8,2), (3,3)}", //48
			"answer",                            //49 TODO: Three points must be plotted on the coordinate plane at ( ̶ 2, 2), (4, ̶ 1), and (8, ̶ 3).
			"\\dfrac{9}{35}",                    //50
		}
		enableCors(&w)
		var grd8mth Grd8mth
		errr := json.NewDecoder(r.Body).Decode(&grd8mth)
		if errr != nil {
			panic(errr)
			return
		}
		rows := [][]string{}
		answered := [50]string{grd8mth.Answer1, grd8mth.Answer2, grd8mth.Answer3, grd8mth.Answer4, grd8mth.Answer5, grd8mth.Answer6, grd8mth.Answer7, grd8mth.Answer8, grd8mth.Answer9, grd8mth.Answer10, grd8mth.Answer11, grd8mth.Answer12, grd8mth.Answer13, grd8mth.Answer14, grd8mth.Answer15, grd8mth.Answer16, grd8mth.Answer17, grd8mth.Answer18, grd8mth.Answer19, grd8mth.Answer20, grd8mth.Answer21, grd8mth.Answer22, grd8mth.Answer23, grd8mth.Answer24, grd8mth.Answer25, grd8mth.Answer26, grd8mth.Answer27, grd8mth.Answer28, grd8mth.Answer29, grd8mth.Answer30, grd8mth.Answer31, grd8mth.Answer32, grd8mth.Answer33, grd8mth.Answer34, grd8mth.Answer35, grd8mth.Answer36, grd8mth.Answer37, grd8mth.Answer38, grd8mth.Answer39, grd8mth.Answer40, grd8mth.Answer41, grd8mth.Answer42, grd8mth.Answer43, grd8mth.Answer44, grd8mth.Answer45, grd8mth.Answer46, grd8mth.Answer47, grd8mth.Answer48, grd8mth.Answer49, grd8mth.Answer50}
		for i := 1; i < 50; i++ {
			if answered[i] == answers[i] {
				rows = append(rows, []string{answered[i], answers[i], "Correct"})
				score++
				question++
			} else {
				rows = append(rows, []string{answered[i], answers[i], "Wrong"})
				question++
			}
		}
		resp := Response{Score: score, Pdf: genpdf("Grade 8 Math", rows, "grd8mth", grd8mth.Name, score, question)}
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			panic(err)
			return
		}
	})
}
