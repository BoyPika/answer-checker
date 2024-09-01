package main

import (
	"encoding/json"
	"net/http"
)

type Worldgeo struct {
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

func worldgeo() {
	http.HandleFunc("/worldgeo", func(w http.ResponseWriter, r *http.Request) {
		score := 0
		question := 0
		answers := [50]string{
			"Subarctic",                              //1
			"Italy",                                  //2
			"Calgary and Durban",                     //3
			"mental mapping",                         //4
			"elevation",                              //5
			"Orientation",                            //6
			"symbols",                                //7
			"Nile",                                   //8
			"Russia",                                 //9
			"cultural regions",                       //10
			"Yugoslavia",                             //11
			"Arctic Ocean",                           //12
			"religious history of Southeast Asia",    //13
			"Persian Gulf",                           //14
			"Section 1 represents economic regions.", //15
			"religion as a dividing force in a region", //16
			"Mediterranean",      //17
			"Sub-Saharan Africa", //18
			"Kremlin",            //19
			"Europe",             //20
			"Toronto",            //21
			"Rwanda",             //22
			"Hinduism",           //23
			"Mountains",          //24
			"Education opportunity pulls Asian students to the United States", //25
			"Population Distribution",            //26
			"Peru",                               //27
			"suburbs",                            //28
			"Japan",                              //29
			"influences agricultural production", //30
			"The number of Chinese people over the age of 65 will increase.", //31
			"cultural diffusion",                       //32
			"locate suitable construction sites",       //33
			"Increased air pollution",                  //34
			"Increased access to trade routes",         //35
			"People became more involved in recycling", //36
			"A dispute over national independence",     //37
			"China",                                    //38
			"Regional district",                        //39
			"Improved health care programs",            //40
			"Skilled workforce",                        //41
			"Internationalization of product assembly", //42
			"Increased global influence",               //43
			"Specialized production",                   //44
			"Greater access to international markets",  //45
			"Environmental Concerns in Box 1",          //46
			"Banking",                                  //47
			"opportunity to lower production costs",    //48
			"Banking in 1",                             //49
			"No nation can employ all of its laborers", //50
		}
		enableCors(&w)
		var worldgeo Worldgeo
		errr := json.NewDecoder(r.Body).Decode(&worldgeo)
		if errr != nil {
			panic(errr)
			return
		}
		rows := [][]string{}
		answered := [50]string{worldgeo.Answer1, worldgeo.Answer2, worldgeo.Answer3, worldgeo.Answer4, worldgeo.Answer5, worldgeo.Answer6, worldgeo.Answer7, worldgeo.Answer8, worldgeo.Answer9, worldgeo.Answer10, worldgeo.Answer11, worldgeo.Answer12, worldgeo.Answer13, worldgeo.Answer14, worldgeo.Answer15, worldgeo.Answer16, worldgeo.Answer17, worldgeo.Answer18, worldgeo.Answer19, worldgeo.Answer20, worldgeo.Answer21, worldgeo.Answer22, worldgeo.Answer23, worldgeo.Answer24, worldgeo.Answer25, worldgeo.Answer26, worldgeo.Answer27, worldgeo.Answer28, worldgeo.Answer29, worldgeo.Answer30, worldgeo.Answer31, worldgeo.Answer32, worldgeo.Answer33, worldgeo.Answer34, worldgeo.Answer35, worldgeo.Answer36, worldgeo.Answer37, worldgeo.Answer38, worldgeo.Answer39, worldgeo.Answer40, worldgeo.Answer41, worldgeo.Answer42, worldgeo.Answer43, worldgeo.Answer44, worldgeo.Answer45, worldgeo.Answer46, worldgeo.Answer47, worldgeo.Answer48, worldgeo.Answer49, worldgeo.Answer50}
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
		resp := Response{Score: score, Pdf: genpdf("World Geography", rows, "worldgeo", worldgeo.Name, score, question)}
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			panic(err)
			return
		}
	})
}
