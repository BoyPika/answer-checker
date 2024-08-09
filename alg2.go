package main

import (
	"encoding/json"
	"net/http"
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

func alg2() {
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

		errr := json.NewDecoder(r.Body).Decode(&alg2)
		if errr != nil {
			panic(errr)
			return
		}
		rows := [][]string{}

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

		genpdf("Algebra 2", rows)

		err := json.NewEncoder(w).Encode(score)
		if err != nil {
			panic(err)
			return
		}
	})
}
