package journey_views

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_views_versions"
)

func init() {
	journey_viewsCmd.AddCommand(journey_views_versions.Cmdjourney_views_versions())
	journey_viewsCmd.Short = utils.GenerateCustomDescription(journey_viewsCmd.Short, journey_views_versions.Description, )
	journey_viewsCmd.Long = journey_viewsCmd.Short
}
