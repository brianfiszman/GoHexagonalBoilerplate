package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/dtos"
)

func GetHeartBeat(rw http.ResponseWriter, r *http.Request) {
	var res dtos.Response = dtos.Response{
		Message: dtos.HELLO_MSG,
		Status:  dtos.STATUS_OK,
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(res)
}
