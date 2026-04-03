package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/arjun-saseendran/banking/service"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zidcode string `json:"zidcode" xml:"zidcode"`
}

type CustomerHandlers struct{
	service service.CustomerService
}


func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{"Maria", "Bridgen", "5094"},
	// 	{"Manjima", "Edukki", "7637"},
	// }
	customers, _ := ch.service.GetAllCustomers()
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)

	}

}
