package rs_controllers

import "github.com/ccutch/view-controller"

type ApplicationController struct {
	*controller.Controller
}

func (this ApplicationController) Home() {
	this.Render("views/homepage.html")
}

// Create a new instance of an application controller
func Application() ApplicationController {
	a := ApplicationController{new(controller.Controller)}
	a.Routes = []controller.Route{
		controller.Route{
			Method:  "GET",
			Path:    "/",
			Handler: a.Home,
		},
	}
	return a
}
