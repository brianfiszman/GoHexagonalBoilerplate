package controllers

import (
	"fmt"
	"net/http"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/config"
	"github.com/go-resty/resty/v2"
)

var restClient resty.Client = *resty.New()

func GetTicketList(rw http.ResponseWriter, r *http.Request) {
	var service_now config.ServiceNowConfig = config.LoadServiceNowConfig()
	res, err := restClient.SetBasicAuth(service_now.USER, service_now.PASS).EnableTrace().GetClient().Get(service_now.API_URL + "/now/table/incident")

	if err != nil {
		http.Error(rw, http.StatusText(401), 401)
	}

	fmt.Fprintf(rw, "%+v", res)
}
