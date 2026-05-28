package journey_views_eventdefinitions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_views_eventdefinitions_activate"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_views_eventdefinitions_external"
)

func init() {
	journey_views_eventdefinitionsCmd.AddCommand(journey_views_eventdefinitions_activate.Cmdjourney_views_eventdefinitions_activate())
	journey_views_eventdefinitionsCmd.AddCommand(journey_views_eventdefinitions_external.Cmdjourney_views_eventdefinitions_external())
	journey_views_eventdefinitionsCmd.Short = utils.GenerateCustomDescription(journey_views_eventdefinitionsCmd.Short, journey_views_eventdefinitions_activate.Description, journey_views_eventdefinitions_external.Description, )
	journey_views_eventdefinitionsCmd.Long = journey_views_eventdefinitionsCmd.Short
}
