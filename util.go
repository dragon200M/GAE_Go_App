package app

import (
	"github.com/julienschmidt/httprouter"

	"fmt"
	"net/http"

	"time"

)

type Czas struct {
	Year string
	Month string
	Day string
	Hour	string
	Minute string
	Second string
}

func MonthtoPolish(month string) string{
	if month == "January"{
		return "Styczeń"
	}

	if month == "February"{
		return "Luty"
	}
	if month == "March"{
		return "Marzec"
	}
	if month == "April"{
		return "Kwiecień"

	}
	if month == "May" {
		return "Maj"
	}
	if month == "June"{
		return "Czerwiec"
	}

	if month == "July"{
		return "Lipiec"
	}
	if month == "August"{
		return "Sierpień"
	}

	if month == "September"{
		return "Wrzesień"
	}

	if month == "October"{
		return "Październik"
	}
	if month == "November"{
		return "Listopad"
	}
	if month == "December"{
		return "Grudzień"
	}

	return month

}


func serveTime(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {


	c:= time.Now()

	t:= c.Format("2006 01 02 15 04 05")


	fmt.Fprint(res,t)

}

