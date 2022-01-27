package services

import (
	"encoding/json"
	"net/http"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/infrastructure/config"
	"github.com/go-chi/jwtauth/v5"
	"github.com/mitchellh/mapstructure"
)

var TokenAuth *jwtauth.JWTAuth = jwtauth.New("HS256", []byte(config.LoadJwtConfig().SECRET), nil)

func CreateJwtToken(r *http.Request) string {
	request := map[string]string{}
	var jsonbody = map[string]interface{}{}

	json.NewDecoder(r.Body).Decode(&request)
	mapstructure.Decode(request, jsonbody)
	_, tokenString, _ := TokenAuth.Encode(jsonbody)

	return tokenString
}
