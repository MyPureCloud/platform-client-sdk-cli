package journey_views

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_views_versions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_views_encodings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_views_eventdefinitions"
)

func init() {
	journey_viewsCmd.AddCommand(journey_views_versions.Cmdjourney_views_versions())
	journey_viewsCmd.AddCommand(journey_views_encodings.Cmdjourney_views_encodings())
	journey_viewsCmd.AddCommand(journey_views_eventdefinitions.Cmdjourney_views_eventdefinitions())
	journey_viewsCmd.Short = utils.GenerateCustomDescription(journey_viewsCmd.Short, journey_views_versions.Description, journey_views_encodings.Description, journey_views_eventdefinitions.Description, )
	journey_viewsCmd.Long = journey_viewsCmd.Short
}
