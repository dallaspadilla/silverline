package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/dallaspadilla/silverline/api/models"
	"github.com/dallaspadilla/silverline/api/responses"
	"github.com/dallaspadilla/silverline/api/utils/formaterror"
	"github.com/gorilla/mux"
)

func (server *Server) CreatePing(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	ping := models.Ping{}
	err = json.Unmarshal(body, &ping)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	ping.Prepare()
	err = ping.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	pingCreated, err := ping.SaveUser(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, pingCreated.PingID))
	responses.JSON(w, http.StatusCreated, pingCreated)
}

func (server *Server) GetPing(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	ping := models.Ping{}
	pingGotten, err := ping.FindUserByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, pingGotten)
}

func (server *Server) GetPings(w http.ResponseWriter, r *http.Request) {

	ping := models.Ping{}

	pings, err := ping.FindAllUsers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, pings)
}
