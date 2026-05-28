package journey_views_eventdefinitions_external

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_views_eventdefinitions_external_changes"
)

func init() {
	journey_views_eventdefinitions_externalCmd.AddCommand(journey_views_eventdefinitions_external_changes.Cmdjourney_views_eventdefinitions_external_changes())
	journey_views_eventdefinitions_externalCmd.Short = utils.GenerateCustomDescription(journey_views_eventdefinitions_externalCmd.Short, journey_views_eventdefinitions_external_changes.Description, )
	journey_views_eventdefinitions_externalCmd.Long = journey_views_eventdefinitions_externalCmd.Short
}
