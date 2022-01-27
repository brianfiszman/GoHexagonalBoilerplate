package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/application/dtos"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/domain/interfaces"
)

type HealthController struct {
	Database interfaces.Database
}

func (h HealthController) IsHealthy(rw http.ResponseWriter, r *http.Request) {
	var res dtos.Response = dtos.Response{
		Message: dtos.HELLO_MSG,
		Status:  dtos.STATUS_OK,
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(res)
}
