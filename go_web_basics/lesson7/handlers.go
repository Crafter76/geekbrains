package main

import (
	"io/ioutil"
	"net/http"
)

// HandleEcho - Эхо функция
// @Summary Эхо функция
// @Param data body string true "Echo body"
// @Description Функция, которая вернет что ей отправили
// @Tags server, echo, system
// @Success 200 {string} string
// @Failure 500 {object} models.ServErr
// @Router /api/v1/echo [post]
func HandleEcho(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err)
		return
	}
	// data = []byte(string(data) + "123")
	w.Write(data)
}

// HandleEcho2 - Эхо функция 2
// @Summary Эхо функция 2
// @Param data body string true "Echo body"
// @Description Функция, которая вернет что ей отправили
// @Tags server, echo
// @Success 200 {string} string
// @Failure 500 {string} string
// @Router /api/v1/echo [put]
func HandleEcho2(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	w.Write(data)
}

// HandleSwaggerJSON - возвращает swagger.json
// @Summary возвращает swagger.json
// @Tags swagger
// @Success 200 {string} string
// @Failure 500 {string} string
// @Router /api/v1/doc/swagger.json [get]
func HandleSwaggerJSON(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./docs/swagger.json")
}
