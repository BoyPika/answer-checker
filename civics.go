package main

import (
	"encoding/json"
	"net/http"
)

type Civics struct {
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
	Name     string `json:"name"`
}

func civics() {
	http.HandleFunc("/civics", func(w http.ResponseWriter, r *http.Request) {
		score := 0
		question := 0
		answers := [32]string{
			"a", //1 TODO: Add include graphics
			"The process of adding amendments is very complex", //2
			"weaknesses within the Articles of Confederation",  //3
			"serving on a jury", //4
			"It addresses local needs and serves the public good.", //5
			"State population",                              //6
			"national policies",                             //7
			"State judicial branch",                         //8
			"Power to influence legislation",                //9
			"Marbury vs Madison",                            //10
			"promote public health",                         //11
			"the national and state governments",            //12
			"Constitution of Virginia",                      //13
			"10 A.M.",                                       //14
			"whether someone is guilty of breaking the law", //15
			"Ranjan",                             //16
			"To help carry out expressed powers", //17
			"Approves the state budget",          //18
			"A health group asks legislators to pass tougher antipollution laws.", //19
			"displayed a positive bias to the candidate from Party X",             //20
			"5th", //21
			"There will not be enough pizza produced.", //22
			"Characteristics That Employers Seek",      //23
			"Free Market — Command",                    //24
			"Guarantee",                                //25
			"A person with a 4-year bachelor's degree and programming experience", //26
			"Income", //27
			"Edward owns a restaurant. He is responsible for all of the restaurant's debts.", //28
			"consumer rights",        //29
			"Lower taxes",            //30
			"global economy",         //31
			"Prohibiting monopolies", //32
		}
		enableCors(&w)
		var civics Civics
		errr := json.NewDecoder(r.Body).Decode(&civics)
		if errr != nil {
			panic(errr)
			return
		}
		rows := [][]string{}
		answered := [32]string{civics.Answer1, civics.Answer2, civics.Answer3, civics.Answer4, civics.Answer5, civics.Answer6, civics.Answer7, civics.Answer8, civics.Answer9, civics.Answer10, civics.Answer11, civics.Answer12, civics.Answer13, civics.Answer14, civics.Answer15, civics.Answer16, civics.Answer17, civics.Answer18, civics.Answer19, civics.Answer20, civics.Answer21, civics.Answer22, civics.Answer23, civics.Answer24, civics.Answer25, civics.Answer26, civics.Answer27, civics.Answer28, civics.Answer29, civics.Answer30, civics.Answer31, civics.Answer32}
		for i := 1; i < 32; i++ {
			if answered[i] == answers[i] {
				rows = append(rows, []string{answered[i], answers[i], "Correct"})
				score++
				question++
			} else {
				rows = append(rows, []string{answered[i], answers[i], "Wrong"})
				question++
			}
		}
		resp := Response{Score: score, Pdf: genpdf("Civics and Economics", rows, "civics", civics.Name, score, question)}
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			panic(err)
			return
		}
	})
}
