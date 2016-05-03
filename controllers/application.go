package controllers

import "github.com/ccutch/go-view-controller"

// Generallized json object
type JSON map[string]interface{}

type ApplicationController struct {
	*controller.Controller
}

// Render homepage
func (this ApplicationController) Home() {
	this.Data = JSON{
		"user": "Connor",
	}
	err := this.Render("homepage")

	if err != nil {
		err = this.RenderError("InternalServer", err)
		if err != nil {
			panic(err)
		}
	}
}

// Create a new instance of an application controller
func Application() ApplicationController {
	a := ApplicationController{new(controller.Controller)}
	//a.Layout = "views/layout.html"
	a.Partials = "views/partials"
	a.Routes = []controller.Route{
		controller.Route{
			Method:  "GET",
			Path:    "/",
			Handler: a.Home,
		},
	}
	return a
}
