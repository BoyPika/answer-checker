package main

import (
	"encoding/json"
	"net/http"
)

type Vastud struct {
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
	Name     string `json:"name"`
}

func vastud() {
	http.HandleFunc("/vastud", func(w http.ResponseWriter, r *http.Request) {
		score := 0
		question := 0
		answers := [34]string{
			"American Indians",                //1
			"digging at archaeological sites", //2
			"Native peoples saw the settlers as invaders taking their land", //3
			"Barn",                       //4
			"move inland",                //5
			"its people opposed slavery", //6
			"4",                          //7
			"Thomas Jefferson",           //8
			"Powhatan",                   //9
			"They rook opposite sides in the American Revolution.", //10
			"Abolitionists",              //11
			"a",                          //12 TODO: Add include graphics
			"Enslaved African Americans", //13
			"Harry F. Byrd, Sr.",         //14
			"Blue Ridge Mountains",       //15
			"James",                      //16
			"Chesapeake Bay",             //17
			"Cherokee",                   //18
			"The Importance of Rivers to the Virginia Settlement", //19
			"a peninsula",                        //20
			"1 and 3",                            //21
			"Capital moved to Williamsburg",      //22
			"English rights to the colonists",    //23
			"had the right to govern themselves", //24
			"Local assemblies",                   //25
			"taken away by Jim Crow laws",        //26
			"make laws for the state",            //27
			"made racial discrimination illegal", //28
			"tourism industry",                   //29
			"Job opportunities",                  //30
			"Bartering in the Colonies",          //31
			"Factories",                          //32
			"To help former slaves",              //33
			"Transportation",                     //34
		}
		enableCors(&w)
		var vastud Vastud
		errr := json.NewDecoder(r.Body).Decode(&vastud)
		if errr != nil {
			panic(errr)
			return
		}
		rows := [][]string{}
		answered := [34]string{vastud.Answer1, vastud.Answer2, vastud.Answer3, vastud.Answer4, vastud.Answer5, vastud.Answer6, vastud.Answer7, vastud.Answer8, vastud.Answer9, vastud.Answer10, vastud.Answer11, vastud.Answer12, vastud.Answer13, vastud.Answer14, vastud.Answer15, vastud.Answer16, vastud.Answer17, vastud.Answer18, vastud.Answer19, vastud.Answer20, vastud.Answer21, vastud.Answer22, vastud.Answer23, vastud.Answer24, vastud.Answer25, vastud.Answer26, vastud.Answer27, vastud.Answer28, vastud.Answer29, vastud.Answer30, vastud.Answer31, vastud.Answer32}
		for i := 1; i < 34; i++ {
			if answered[i] == answers[i] {
				rows = append(rows, []string{answered[i], answers[i], "Correct"})
				score++
				question++
			} else {
				rows = append(rows, []string{answered[i], answers[i], "Wrong"})
				question++
			}
		}
		resp := Response{Score: score, Pdf: genpdf("Virginia Studies", rows, "vastud", vastud.Name, score, question)}
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			panic(err)
			return
		}
	})
}
