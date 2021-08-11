package controllers

import (
	"fmt"
	"net/http"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/services"
)


func Auth(rw http.ResponseWriter, r *http.Request){
	var jwt = services.CreateJwtToken(r)
	
	fmt.Fprintf(rw, "%+v", jwt)
}