package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/brianfiszman/GoHexagonalBoilerplate/pkg/application/dtos"
	"github.com/brianfiszman/GoHexagonalBoilerplate/pkg/domain/services"
	"github.com/sirupsen/logrus"
)

type HealthController struct {
	HealthService services.HealthService
}

func (h HealthController) IsHealthy(rw http.ResponseWriter, r *http.Request) {
	var res dtos.Response = dtos.Response{
		Message: dtos.HELLO_MSG,
		Status:  dtos.STATUS_OK,
	}

	err := h.HealthService.GetHealthiness()

	if err != nil {
		logrus.Error(err)

		res.Message = err.Error()
		res.Status = http.StatusInternalServerError
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(res.Status)
	json.NewEncoder(rw).Encode(res)
}
