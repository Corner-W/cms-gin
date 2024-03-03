package router

import (

	"cms/router/cms"
)

type Group struct {

	
	Cms cms.RouterGroup
}

var GroupApp = new(Group)
