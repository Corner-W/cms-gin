package v1

import (
	"cms/api/v1/cms"

)

type ApiGroup struct {

	CmsGroup cms.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
