package main

import (
	"encoding/json"
	"net/http"
)

type Chem struct {
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
	Name     string `json:"name"`
}

func chem() {
	http.HandleFunc("/chem", func(w http.ResponseWriter, r *http.Request) {
		score := 0
		question := 0
		answers := [46]string{
			"1\\%",                              //1
			"\\textnormal{a temperature probe}", //2
			"\\textnormal{Double-replacement}",  //3
			"(2-)",                              //4 //TODO(MULTIPLE ANSWERS) (2-, -2)
			"\\textnormal{0.4 M}",               //5
			"39.3",                              //6
			"( .806)g/mL",                       //7 //TODO (MULTIPLE ANSWERS) (0.806 g/mL, .806g/mL)
			"\\textnormal{Phosphorus (P)}",      //8
			"\\textnormal{Single-replacement}",  //9
			"\\textnormal{the positive ion}",    //10
			"\\textnormal{Filtering}",           //11
			"\\textnormal{Store all solutions in brown bottles}", //12
			"\\textnormal{Number of the solute particles}",       //13
			"\\textnormal{Ammonium nitrate}",                     //14
			"(1)",                                                //15 //TODO (MULTIPLE ANSWERS) (1, 1., 1.3, 1.32, 1.33, 1.4, 2, 2.) (Yes 2 and 2. are different, welcome to chemistry)
			"\\textnormal{Phosphorus (P)}",                       //16
		}

		var chem Chem

		errr := json.NewDecoder(r.Body).Decode(&chem)
		if errr != nil {
			panic(errr)
			return
		}
		enableCors(&w)
		rows := [][]string{}
		answered := [46]string{chem.Answer1, chem.Answer2, chem.Answer3, chem.Answer4, chem.Answer5, chem.Answer6, chem.Answer7, chem.Answer8, chem.Answer9, chem.Answer10, chem.Answer11, chem.Answer12, chem.Answer13, chem.Answer14, chem.Answer15, chem.Answer16, chem.Answer17, chem.Answer18, chem.Answer19, chem.Answer20, chem.Answer21, chem.Answer22, chem.Answer23, chem.Answer24, chem.Answer25, chem.Answer26, chem.Answer27, chem.Answer28, chem.Answer29, chem.Answer30, chem.Answer31, chem.Answer32, chem.Answer33, chem.Answer34, chem.Answer35, chem.Answer36, chem.Answer37, chem.Answer38, chem.Answer39, chem.Answer40, chem.Answer41, chem.Answer42, chem.Answer43, chem.Answer44, chem.Answer45, chem.Answer46}
		for i := 1; i < 46; i++ {
			if answered[i] == answers[i] {
				rows = append(rows, []string{answered[i], answers[i], "Correct"})
				score++
				question++
			} else {
				rows = append(rows, []string{answered[i], answers[i], "Wrong"})
				question++
			}
		}
		resp := Response{Score: score, Pdf: genpdf("Chemistry", rows, "chemistry", chem.Name, score, question)}
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			panic(err)
			return
		}
	})
}
