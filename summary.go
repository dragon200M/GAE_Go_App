package app

import (
	"net/http"


	"strconv"
)

func summary(req *http.Request, usr *User) []Summary{

	exp,_ :=getExpenses(req,usr)


	b := make(map[string][]Expenses)

	for _,v :=range exp {
		b[v.Month.String()] = append(b[v.Month.String()],v)

	}

	var summary2 []Summary

	for k,v :=range b {
		m :=map[string]float64{}
		sumMonth := 0.0
		for i:=0;i<len(v);i++{
			m[v[i].Category]+=v[i].Amount
			sumMonth += v[i].Amount
		}
		var catSum []CatSum

		for c,v :=range m {
			f :=CatSum{Name:c, Sum:strconv.FormatFloat(v,'f',2,64)}
			catSum = append(catSum,f)

		}
		summary := Summary{Month:MonthtoPolish(k), CatSum:catSum, MonthSum:strconv.FormatFloat(sumMonth,'f',2,64)}
		summary2 = append(summary2,summary)

	}


	return summary2
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
