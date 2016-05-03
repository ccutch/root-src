package rs_controllers

import "github.com/ccutch/go-view-controller"

type ApplicationController struct {
	*controller.Controller
}

// Render homepage
func (this ApplicationController) Home() {
	this.Data = struct {
		Title string
		Name  string
	}{Title: "HOMEPAGE", Name: "Connor"}
	this.Render("homepage")
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
