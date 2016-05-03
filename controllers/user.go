package controllers

import (
	"net/http"

	"github.com/ccutch/go-view-controller"
	"github.com/ccutch/root-src/models"
)

type UserController struct {
	*controller.Controller
}

func (u UserController) Invoke() {
	http.Redirect(u.Res, u.Req, models.GetOauthUrl(), 302)
}

func User() UserController {
	u := UserController{new(controller.Controller)}
	u.Routes = []controller.Route{
		controller.Route{
			Method:  "GET",
			Path:    "/invoke",
			Handler: u.Invoke,
		},
	}
	return u
}
