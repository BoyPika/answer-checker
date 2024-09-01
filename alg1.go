package main

import (
	"encoding/json"
	"net/http"
)

type Alg1 struct {
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
	Name     string `json:"name"`
}

func alg1() {
	http.HandleFunc("/alg1", func(w http.ResponseWriter, r *http.Request) {
		score := 0
		question := 0
		answers := [44]string{
			"\\dfrac{1}{2}n-4",      //1
			"x-3",                   //2
			"B,C",                   //3
			"x-2y",                  //4
			"2\\sqrt[3]{48}",        //5
			"8\\sqrt{2}",            //6
			"3-2n",                  //7
			"n-6",                   //8
			"\\dfrac{a^{9}}{b^{4}}", //9
			"-18",                   //10
			"a",                     //11 TODO: Add include graphics
			"x=\\dfrac{180-y}{2}",   //12
			"y=-\\dfrac{1}{2}x+2",   //13
			"\\set{{x+y}\\leq{-2}}{{2x-3y}\\leq{-9}}", //14
			"(1,5)",                                //15
			"answer",                               //16 TODO: Fix the frontend for this
			"4",                                    //17
			"y=\\dfrac{1}{3}x-4",                   //18
			"answer",                               //19 TODO: Fix the frontend for this
			"(-17)",                                //20
			"-5 \\textnormal{ and } \\dfrac{4}{3}", //21
			"-\\dfrac{1}{6}",                       //22
			"6-y",                                  //23 Technically also y+-6
			"(3,-7)",                               //24
			"5",                                    //25
			"d",                                    //26 TODO: Add include graphics
			"",                                     //27 TODO: (2, ̶ 8), (0,0), (1,4), or (2,8)
			"b",                                    //28 TODO: Add include graphics
			"y=1.1x^{2}+4.2x-4.9",                  //29
			"y=x^{2}+1",                            //30
			"2000 = An",                            //31
			"\\textnormal{All real numbers greater than or equal to -4}", //32
			"", //33 TODO: Answers must be placed in the correct order from left to right: Set 3; Set 1; Set 2
			"\\textnormal{The range of the data for Intersection 2 is twice the range of the data for Intersection 1}", //34
			"k(x)=x^{2}+11x+24", //35
			"b",                 //36 TODO: Fix the frontend for this
			"y=x^{2}+3x-28",     //37
			"-6",                //38
			"B",                 //39 TODO: Add include graphics
			"10",                //40
			"4",                 //41
			"170",               //42
			"-1.83",             //43
			"a",                 //44 TODO: Fix the frontend for this
		}

		enableCors(&w)

		var alg1 Alg1

		errr := json.NewDecoder(r.Body).Decode(&alg1)
		if errr != nil {
			panic(errr)
			return
		}
		rows := [][]string{}

		answered := [44]string{alg1.Answer1, alg1.Answer2, alg1.Answer3, alg1.Answer4, alg1.Answer5, alg1.Answer6, alg1.Answer7, alg1.Answer8, alg1.Answer9, alg1.Answer10, alg1.Answer11, alg1.Answer12, alg1.Answer13, alg1.Answer14, alg1.Answer15, alg1.Answer16, alg1.Answer17, alg1.Answer18, alg1.Answer19, alg1.Answer20, alg1.Answer21, alg1.Answer22, alg1.Answer23, alg1.Answer24, alg1.Answer25, alg1.Answer26, alg1.Answer27, alg1.Answer28, alg1.Answer29, alg1.Answer30, alg1.Answer31, alg1.Answer32, alg1.Answer33, alg1.Answer34, alg1.Answer35, alg1.Answer36, alg1.Answer37, alg1.Answer38, alg1.Answer39, alg1.Answer40, alg1.Answer41, alg1.Answer42, alg1.Answer43, alg1.Answer44}
		for i := 1; i < 44; i++ {
			if answered[i] == answers[i] {
				rows = append(rows, []string{answered[i], answers[i], "Correct"})
				score++
				question++
			} else {
				rows = append(rows, []string{answered[i], answers[i], "Wrong"})
				question++
			}
		}
		resp := Response{Score: score, Pdf: genpdf("Algebra 1", rows, "algebra1", alg1.Name, score, question)}
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			panic(err)
			return
		}
	})
}
