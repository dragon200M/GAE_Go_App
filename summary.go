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




