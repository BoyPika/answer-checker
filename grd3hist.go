package main

import (
	"encoding/json"
	"net/http"
)

type Grd3hist struct {
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
	Name     string `json:"name"`
}

func grd3hist() {
	http.HandleFunc("/grd3hist", func(w http.ResponseWriter, r *http.Request) {
		score := 0
		question := 0
		answers := [33]string{
			"Powhatan",              //1
			"Used written language", //2
			"James River",           //3
			"People elect a small group to make the laws.", //4
			"d",                                      //5 TODO: Add include graphics
			"Betsy Ross",                             //6
			"peanuts",                                //7
			"France",                                 //8
			"Olympic games",                          //9
			"Mali",                                   //10
			"2",                                      //11
			"fishermen",                              //12
			"Rome and Greece",                        //13
			"Virginia",                               //14
			"A4",                                     //15
			"sunny and warm in each season",          //16
			"raise more crops",                       //17
			"North",                                  //18
			"Pueblo",                                 //19
			"Africa",                                 //20
			"2 and 4",                                //21
			"c",                                      //22 TODO: Add include graphics
			"1",                                      //23
			"Saving",                                 //24
			"Mr. Carson who is working as a farmer.", //25
			"a capital resource",                     //26
			"Memorial Day",                           //27
			"A firefighter putting out a house fire", //28
			"Voting in elections",                    //29
			"lawyer",                                 //30
			"had disabilities",                       //31
			"choose state leaders",                   //32
			"improving the lives of American farm workers", //33
		}
		enableCors(&w)
		var grd3hist Grd3hist
		errr := json.NewDecoder(r.Body).Decode(&grd3hist)
		if errr != nil {
			panic(errr)
			return
		}
		rows := [][]string{}
		answered := [33]string{grd3hist.Answer1, grd3hist.Answer2, grd3hist.Answer3, grd3hist.Answer4, grd3hist.Answer5, grd3hist.Answer6, grd3hist.Answer7, grd3hist.Answer8, grd3hist.Answer9, grd3hist.Answer10, grd3hist.Answer11, grd3hist.Answer12, grd3hist.Answer13, grd3hist.Answer14, grd3hist.Answer15, grd3hist.Answer16, grd3hist.Answer17, grd3hist.Answer18, grd3hist.Answer19, grd3hist.Answer20, grd3hist.Answer21, grd3hist.Answer22, grd3hist.Answer23, grd3hist.Answer24, grd3hist.Answer25, grd3hist.Answer26, grd3hist.Answer27, grd3hist.Answer28, grd3hist.Answer29, grd3hist.Answer30, grd3hist.Answer31, grd3hist.Answer32}
		for i := 1; i < 33; i++ {
			if answered[i] == answers[i] {
				rows = append(rows, []string{answered[i], answers[i], "Correct"})
				score++
				question++
			} else {
				rows = append(rows, []string{answered[i], answers[i], "Wrong"})
				question++
			}
		}
		resp := Response{Score: score, Pdf: genpdf("Virginia Studies", rows, "grd3hist", grd3hist.Name, score, question)}
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			panic(err)
			return
		}
	})
}
