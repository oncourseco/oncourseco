package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/unrealities/warning-track/models"
	"google.golang.org/appengine"
)

func oncourse(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	info := []models.Info{}

	c := appengine.NewContext(r)
}
